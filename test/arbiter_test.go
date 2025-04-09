package test

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/client"
	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/auth"
	"github.com/openshift/assisted-service/pkg/requestid"
	"github.com/openshift/assisted-service/subsystem/utils_test"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func registerInfraEnv(clusterID *strfmt.UUID, imageType models.ImageType, cpuArch, ocpVersion, pullSecret string) *models.InfraEnv {
	request, err := utils_test.TestContext.UserBMClient.Installer.RegisterInfraEnv(context.Background(), &installer.RegisterInfraEnvParams{
		InfraenvCreateParams: &models.InfraEnvCreateParams{
			Name:             swag.String("test-infra-env"),
			OpenshiftVersion: ocpVersion,
			PullSecret:       swag.String(pullSecret),
			SSHAuthorizedKey: swag.String(utils_test.SshPublicKey),
			ImageType:        imageType,
			ClusterID:        clusterID,
			CPUArchitecture:  cpuArch,
		},
	})

	if err != nil {
		log.Info(err)
		panic(err)
	}
	return request.GetPayload()
}

func registerHostsAndSetRoles(clusterID, infraenvID strfmt.UUID, clusterName string, baseDNSDomain string, machineCIDR string) []*models.Host {
	ctx := context.Background()
	hosts := make([]*models.Host, 0)

	ips := hostutil.GenerateIPv4Addresses(3, machineCIDR)
	for i := 0; i < 3; i++ {
		hostname := fmt.Sprintf("h%d", i)
		inventory := utils_test.GetDefaultInventory(ips[i])
		inventory.Interfaces[0].MacAddress = fmt.Sprintf("e6:53:3d:a7:77:b%d", i)
		host := utils_test.TestContext.RegisterNodeWithInventory(ctx, infraenvID, hostname, ips[i], inventory)
		var role models.HostRole
		if i < 2 {
			role = models.HostRoleMaster
		} else {
			role = models.HostRoleArbiter
		}
		_, err := utils_test.TestContext.UserBMClient.Installer.V2UpdateHost(ctx, &installer.V2UpdateHostParams{
			HostUpdateParams: &models.HostUpdateParams{
				HostRole: swag.String(string(role)),
			},
			HostID:     *host.ID,
			InfraEnvID: infraenvID,
		})
		if err != nil {
			log.Info(err)
			panic(err)
		}
		hosts = append(hosts, host)
	}
	for _, host := range hosts {
		utils_test.TestContext.GenerateDomainResolution(ctx, host, clusterName, baseDNSDomain)
		utils_test.TestContext.GenerateCommonDomainReply(ctx, host, clusterName, baseDNSDomain)
	}
	generateFullMeshConnectivity(ctx, ips[0], hosts...)
	cluster := utils_test.TestContext.GetCluster(clusterID)
	if cluster.DiskEncryption != nil && swag.StringValue(cluster.DiskEncryption.Mode) == models.DiskEncryptionModeTang {
		utils_test.TestContext.GenerateTangPostStepReply(ctx, true, hosts...)
	}

	if !swag.BoolValue(cluster.UserManagedNetworking) {
		_, err := utils_test.TestContext.UserBMClient.Installer.V2UpdateCluster(ctx, &installer.V2UpdateClusterParams{
			ClusterUpdateParams: &models.V2ClusterUpdateParams{
				VipDhcpAllocation: swag.Bool(false),
				APIVips:           []*models.APIVip{},
				IngressVips:       []*models.IngressVip{},
			},
			ClusterID: clusterID,
		})
		if err != nil {
			log.Info(err)
			panic(err)
		}
		apiVip := "1.2.3.8"
		ingressVip := "1.2.3.9"
		_, err = utils_test.TestContext.UserBMClient.Installer.V2UpdateCluster(ctx, &installer.V2UpdateClusterParams{
			ClusterUpdateParams: &models.V2ClusterUpdateParams{
				APIVips:     []*models.APIVip{{IP: models.IP(apiVip), ClusterID: clusterID}},
				IngressVips: []*models.IngressVip{{IP: models.IP(ingressVip), ClusterID: clusterID}},
			},
			ClusterID: clusterID,
		})

		if err != nil {
			log.Info(err)
			panic(err)
		}
	}

	waitForHostState(ctx, models.HostStatusKnown, utils_test.DefaultWaitForHostStateTimeout, hosts...)
	waitForClusterState(ctx, clusterID, models.ClusterStatusReady, 60*time.Second, utils_test.ClusterReadyStateInfo)

	return hosts
}

func generateFullMeshConnectivity(ctx context.Context, startCIDR string, hosts ...*models.Host) {

	ip, _, err := net.ParseCIDR(startCIDR)
	if err != nil {
		log.Info(err)
		panic(err)
	}
	hostToAddr := make(map[strfmt.UUID]string)

	for _, h := range hosts {
		hostToAddr[*h.ID] = ip.String()
		common.IncrementIP(ip)
	}

	var connectivityReport models.ConnectivityReport
	for _, h := range hosts {

		l2Connectivity := make([]*models.L2Connectivity, 0)
		l3Connectivity := make([]*models.L3Connectivity, 0)
		for id, addr := range hostToAddr {

			if id != *h.ID {
				continue
			}

			l2Connectivity = append(l2Connectivity, &models.L2Connectivity{
				RemoteIPAddress: addr,
				Successful:      true,
			})
			l3Connectivity = append(l3Connectivity, &models.L3Connectivity{
				RemoteIPAddress: addr,
				Successful:      true,
			})
		}

		connectivityReport.RemoteHosts = append(connectivityReport.RemoteHosts, &models.ConnectivityRemoteHost{
			HostID:         *h.ID,
			L2Connectivity: l2Connectivity,
			L3Connectivity: l3Connectivity,
		})
	}

	for _, h := range hosts {
		generateConnectivityPostStepReply(ctx, h, &connectivityReport)
	}
}

func generateConnectivityPostStepReply(ctx context.Context, h *models.Host, connectivityReport *models.ConnectivityReport) {
	fa, err := json.Marshal(connectivityReport)
	if err != nil {
		log.Info(err)
		panic(err)
	}
	_, err = utils_test.TestContext.AgentBMClient.Installer.V2PostStepReply(ctx, &installer.V2PostStepReplyParams{
		InfraEnvID: h.InfraEnvID,
		HostID:     *h.ID,
		Reply: &models.StepReply{
			ExitCode: 0,
			Output:   string(fa),
			StepID:   string(models.StepTypeConnectivityCheck),
			StepType: models.StepTypeConnectivityCheck,
		},
	})
	if err != nil {
		log.Info(err)
		panic(err)
	}
}

func waitForHostState(ctx context.Context, state string, timeout time.Duration, hosts ...*models.Host) {
	waitForHost := func(host *models.Host) error {
		log.Infof("Waiting for host %s state %s", host.ID.String(), state)
		var (
			lastState      string
			lastStatusInfo string
			success        bool
		)

		for start, successInRow := time.Now(), 0; time.Since(start) < timeout; {
			success, lastState, lastStatusInfo = isHostInState(ctx, host.InfraEnvID, *host.ID, state)

			if success {
				successInRow++
			} else {
				successInRow = 0
			}

			// Wait for host state to be consistent
			if successInRow >= utils_test.MinSuccessesInRow {
				log.Infof("host %s has status %s", host.ID.String(), state)
				return nil
			}

			time.Sleep(time.Second)
		}

		if lastState != state {
			return fmt.Errorf("Host %s wasn't in state %s for %d times in a row. Actual %s (%s)",
				host.ID.String(), state, utils_test.MinSuccessesInRow, lastState, lastStatusInfo)
		}

		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	errs, _ := errgroup.WithContext(ctx)
	defer cancel()

	for idx := range hosts {
		host := hosts[idx]
		errs.Go(func() error {
			return waitForHost(host)
		})
	}

	err := errs.Wait()
	if err != nil {
		log.Info(err)
		panic(err)
	}
}

func isHostInState(ctx context.Context, infraEnvID strfmt.UUID, hostID strfmt.UUID, state string) (bool, string, string) {
	rep, err := utils_test.TestContext.UserBMClient.Installer.V2GetHost(ctx, &installer.V2GetHostParams{
		InfraEnvID: infraEnvID,
		HostID:     hostID,
	})
	if err != nil {
		log.Info(err)
		panic(err)
	}
	h := rep.GetPayload()
	return swag.StringValue(h.Status) == state, swag.StringValue(h.Status), swag.StringValue(h.StatusInfo)
}

func waitForClusterState(ctx context.Context, clusterID strfmt.UUID, state string, timeout time.Duration, stateInfo string) {
	log.Infof("Waiting for cluster %s status %s", clusterID, state)
	var (
		lastState      string
		lastStatusInfo string
		success        bool
	)

	for start, successInRow := time.Now(), 0; time.Since(start) < timeout; {
		success, lastState, lastStatusInfo = isClusterInState(ctx, clusterID, state, stateInfo)

		if success {
			successInRow++
		} else {
			successInRow = 0
		}

		// Wait for cluster state to be consistent
		if successInRow >= utils_test.MinSuccessesInRow {
			log.Infof("cluster %s has status %s", clusterID, state)
			return
		}

		time.Sleep(time.Second)
	}

	if lastState != state {
		err := fmt.Sprintf("Cluster %s wasn't in state %s for %d times in a row. Actual %s (%s)",
			clusterID, state, utils_test.MinSuccessesInRow, lastState, lastStatusInfo)
		log.Info(err)
		panic(err)
	}
}

func isClusterInState(ctx context.Context, clusterID strfmt.UUID, state, stateInfo string) (bool, string, string) {
	rep, err := utils_test.TestContext.UserBMClient.Installer.V2GetCluster(ctx, &installer.V2GetClusterParams{ClusterID: clusterID})
	if err != nil {
		log.Info(err)
		panic(err)
	}
	c := rep.GetPayload()
	if swag.StringValue(c.Status) == state {
		return stateInfo == utils_test.IgnoreStateInfo ||
			swag.StringValue(c.StatusInfo) == stateInfo, swag.StringValue(c.Status), swag.StringValue(c.StatusInfo)
	}
	if swag.StringValue(c.Status) == "error" {
		log.Info("c.status is error")
		panic("c.status is error")
	}

	return false, swag.StringValue(c.Status), swag.StringValue(c.StatusInfo)
}

func createBmInventoryClient(pullSecretToken string) (*client.AssistedInstall, error) {
	clientConfig := client.Config{}
	u, err := url.Parse("http://localhost:8090/api/assisted-install")
	if err != nil {
		return nil, errors.Wrap(err, "Failed parsing inventory URL")
	}
	clientConfig.URL = u

	transport := requestid.Transport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	minDelay := time.Duration(2) * time.Second
	maxDelay := time.Duration(10) * time.Second
	retries := 3

	// This function transforms a regular delay function into one that does the same thing but
	// also writes a log message indicating that the request will be retried.
	delayLog := func(delayFn rehttp.DelayFn) rehttp.DelayFn {
		return func(attempt rehttp.Attempt) time.Duration {
			delay := delayFn(attempt)
			fields := logrus.Fields{
				"method":  attempt.Request.Method,
				"url":     attempt.Request.URL,
				"error":   attempt.Error,
				"attempt": fmt.Sprintf("%d of %d", attempt.Index+1, retries+1),
				"delay":   delay,
			}
			if attempt.Response != nil {
				fields["status"] = attempt.Response.StatusCode
			}
			logrus.WithFields(fields).Info("Request will be retried")
			return delay
		}
	}

	// Add retry settings
	tr := rehttp.NewTransport(
		transport,
		rehttp.RetryAll(
			rehttp.RetryMaxRetries(retries),
			rehttp.RetryAny(
				rehttp.RetryTemporaryErr(),
				rehttp.RetryStatuses(
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout,
				),
			),
		),
		delayLog(rehttp.ExpJitterDelay(minDelay, maxDelay)),
	)

	clientConfig.Transport = tr

	clientConfig.AuthInfo = auth.AgentAuthHeaderWriter(pullSecretToken)
	bmInventory := client.New(clientConfig)
	rtctransport := bmInventory.Transport.(*rtclient.Runtime)
	rtctransport.Consumers[runtime.HTMLMime] = HTMLConsumer()
	return bmInventory, nil
}

func HTMLConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if reader == nil {
			return errors.New("HTMLConsumer requires a reader") // early exit
		}

		//read the response body
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(reader)
		if err != nil {
			return err
		}
		b := buf.Bytes()
		msg := string(b)

		//handle empty response body
		if len(b) == 0 {
			return nil
		}

		t := reflect.TypeOf(data)
		if data == nil || t.Kind() != reflect.Ptr {
			return fmt.Errorf("data should be a non nil pointer")
		}

		switch dt := data.(type) {
		case string:
			v := reflect.Indirect(reflect.ValueOf(data))
			v.SetString(msg)
		case encoding.TextUnmarshaler:
			return dt.UnmarshalText(b)
		case *models.Error:
			dt.Reason = swag.String(msg)
		case *models.InfraError:
			dt.Message = swag.String(msg)
		default:
			return fmt.Errorf("%+v (%T) is not supported by the Agent's Custom Consumer", data, data)
		}

		return nil
	})
}

func installCluster(clusterID strfmt.UUID) *models.Cluster {
	ctx := context.Background()
	reply, err := utils_test.TestContext.UserBMClient.Installer.V2InstallCluster(ctx, &installer.V2InstallClusterParams{ClusterID: clusterID})
	Expect(err).NotTo(HaveOccurred())
	c := reply.GetPayload()
	Expect(*c.Status).Should(Equal(models.ClusterStatusPreparingForInstallation))
	utils_test.TestContext.GenerateEssentialPrepareForInstallationSteps(ctx, c.Hosts...)

	waitForClusterState(ctx, clusterID, models.ClusterStatusInstalling,
		180*time.Second, "Installation in progress")

	waitForHostState(ctx, models.HostStatusInstalling, utils_test.DefaultWaitForHostStateTimeout, c.Hosts...)

	rep, err := utils_test.TestContext.UserBMClient.Installer.V2GetCluster(ctx, &installer.V2GetClusterParams{ClusterID: clusterID})
	Expect(err).NotTo(HaveOccurred())
	c = rep.GetPayload()
	Expect(c).NotTo(BeNil())

	return c
}

func setClusterAsInstalling(ctx context.Context, clusterID strfmt.UUID) {
	c := installCluster(clusterID)
	Expect(swag.StringValue(c.Status)).Should(Equal("installing"))
	Expect(swag.StringValue(c.StatusInfo)).Should(Equal("Installation in progress"))

	for _, host := range c.Hosts {
		Expect(swag.StringValue(host.Status)).Should(Equal("installing"))
	}
}

func setClusterAsFinalizing(ctx context.Context, clusterID strfmt.UUID) {
	setClusterAsInstalling(ctx, clusterID)
	c := utils_test.TestContext.GetCluster(clusterID)

	for _, host := range c.Hosts {
		utils_test.TestContext.UpdateProgress(*host.ID, host.InfraEnvID, models.HostStageDone)
	}

	waitForClusterState(ctx, clusterID, models.ClusterStatusFinalizing, utils_test.DefaultWaitForClusterStateTimeout, utils_test.ClusterFinalizingStateInfo)
}

func completeInstallation(client *client.AssistedInstall, clusterID strfmt.UUID) {
	ctx := context.Background()
	rep, err := client.Installer.V2GetCluster(ctx, &installer.V2GetClusterParams{ClusterID: clusterID})
	Expect(err).NotTo(HaveOccurred())

	status := models.OperatorStatusAvailable

	Eventually(func() error {
		_, err = utils_test.TestContext.AgentBMClient.Installer.V2UploadClusterIngressCert(ctx, &installer.V2UploadClusterIngressCertParams{
			ClusterID:         clusterID,
			IngressCertParams: models.IngressCertParams(utils_test.IngressCa),
		})
		return err
	}, "10s", "2s").Should(BeNil())

	for _, operator := range rep.Payload.MonitoredOperators {
		if operator.OperatorType != models.OperatorTypeBuiltin {
			continue
		}

		utils_test.TestContext.V2ReportMonitoredOperatorStatus(ctx, clusterID, operator.Name, status, "")
	}
}

func completeInstallationAndVerify(ctx context.Context, client *client.AssistedInstall, clusterID strfmt.UUID) {
	completeInstallation(client, clusterID)
	waitForClusterState(ctx, clusterID, models.ClusterStatusInstalled, utils_test.DefaultWaitForClusterStateTimeout, utils_test.IgnoreStateInfo)
}

const (
	pollDefaultInterval = 1 * time.Millisecond
	pollDefaultTimeout  = 30 * time.Second
)

var _ = Describe("Create TNA", func() {
	It("test", func() {
		var (
			ctx              = context.Background()
			cluster          *models.Cluster
			infraEnvID       *strfmt.UUID
			openshiftVersion = "4.19.0-0.nightly-2025-03-06-121124"
			clusterCIDR      = "10.128.0.0/14"
			serviceCIDR      = "172.30.0.0/16"
			machineCIDR      = "1.2.3.0/24"
			pullSecret       = ""
		)

		bmClient, err := createBmInventoryClient(pullSecret)
		if err != nil {
			log.Info(err)
			panic(err)
		}
		utils_test.TestContext = utils_test.NewSubsystemTestContext(
			log.New(),
			nil,
			bmClient,
			bmClient,
			bmClient,
			bmClient,
			bmClient,
			bmClient,
			bmClient,
			bmClient,
			pollDefaultInterval,
			pollDefaultTimeout,
			"4.14.0",
		)

		registerClusterReply, err := utils_test.TestContext.UserBMClient.Installer.V2RegisterCluster(ctx, &installer.V2RegisterClusterParams{
			NewClusterParams: &models.ClusterCreateParams{
				BaseDNSDomain:        "example.com",
				ClusterNetworks:      []*models.ClusterNetwork{{Cidr: models.Subnet(clusterCIDR), HostPrefix: 23}},
				ServiceNetworks:      []*models.ServiceNetwork{{Cidr: models.Subnet(serviceCIDR)}},
				Name:                 swag.String("test-cluster"),
				OpenshiftVersion:     swag.String(openshiftVersion),
				PullSecret:           swag.String(pullSecret),
				SSHPublicKey:         utils_test.SshPublicKey,
				HighAvailabilityMode: swag.String("TNA"),
				ControlPlaneCount:    swag.Int64(2),
			},
		})
		if err != nil {
			log.Info(err)
			panic(err)
		}
		cluster = registerClusterReply.GetPayload()
		log.Infof("Register cluster %s", cluster.ID.String())
		infraEnvID = registerInfraEnv(cluster.ID, models.ImageTypeMinimalIso, "x86_64", openshiftVersion, pullSecret).ID
		clusterID := *cluster.ID
		registerHostsAndSetRoles(clusterID, *infraEnvID, cluster.Name, cluster.BaseDNSDomain, machineCIDR)
		setClusterAsFinalizing(ctx, clusterID)
		completeInstallationAndVerify(ctx, utils_test.TestContext.AgentBMClient, clusterID)
	})
})

func TestArbiter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arbiter test")
}
