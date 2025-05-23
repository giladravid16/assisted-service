package network

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/go-openapi/swag"
	"github.com/kelseyhightower/envconfig"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	manifestsapi "github.com/openshift/assisted-service/internal/manifests/api"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/tang"
	operations "github.com/openshift/assisted-service/restapi/operations/manifests"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

//go:generate mockgen -source=manifests_generator.go -package=network -destination=mock_manifests_generator.go

type ManifestsGeneratorAPI interface {
	AddChronyManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	AddDnsmasqForSingleNode(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	AddTelemeterManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	AddSchedulableMastersManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	AddDiskEncryptionManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	AddNicReapply(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error
	IsSNODNSMasqEnabled() bool
}

type Config struct {
	ServiceBaseURL          string `envconfig:"SERVICE_BASE_URL"`
	EnableSingleNodeDnsmasq bool   `envconfig:"ENABLE_SINGLE_NODE_DNSMASQ" default:"false"`
}

type ManifestsGenerator struct {
	manifestsApi manifestsapi.ManifestsAPI
	Config       Config
	DB           *gorm.DB
}

func NewManifestsGenerator(manifestsApi manifestsapi.ManifestsAPI, config Config, db *gorm.DB) *ManifestsGenerator {
	return &ManifestsGenerator{
		manifestsApi: manifestsApi,
		Config:       config,
		DB:           db,
	}
}

const defaultChronyConf = `
pool 0.rhel.pool.ntp.org iburst
driftfile /var/lib/chrony/drift
makestep 1.0 3
rtcsync
logdir /var/log/chrony`

const ntpMachineConfigManifest = `
apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  labels:
    machineconfiguration.openshift.io/role: {{.ROLE}}
  name: 50-{{.ROLE}}s-chrony-configuration
spec:
  config:
    ignition:
      version: 3.1.0
    storage:
      files:
      - contents:
          source: data:text/plain;charset=utf-8;base64,{{.CHRONY_CONTENT}}
        mode: 420
        path: /etc/chrony.conf
        overwrite: true
`

const snoDnsmasqConf = `#!/usr/bin/env bash

# In order to override cluster domain please provide this file with the following params:
# SNO_CLUSTER_NAME_OVERRIDE=<new cluster name>
# SNO_BASE_DOMAIN_OVERRIDE=<your new base domain>
# SNO_DNSMASQ_IP_OVERRIDE=<new ip>
if [ -f /etc/default/sno_dnsmasq_configuration_overrides ]; then
    source /etc/default/sno_dnsmasq_configuration_overrides
fi

HOST_IP=${SNO_DNSMASQ_IP_OVERRIDE:-"{{.HOST_IP}}"}
CLUSTER_NAME=${SNO_CLUSTER_NAME_OVERRIDE:-"{{.CLUSTER_NAME}}"}
BASE_DOMAIN=${SNO_BASE_DOMAIN_OVERRIDE:-"{{.DNS_DOMAIN}}"}
CLUSTER_FULL_DOMAIN="${CLUSTER_NAME}.${BASE_DOMAIN}"

cat << EOF > /etc/dnsmasq.d/single-node.conf
address=/apps.${CLUSTER_FULL_DOMAIN}/${HOST_IP}
address=/api-int.${CLUSTER_FULL_DOMAIN}/${HOST_IP}
address=/api.${CLUSTER_FULL_DOMAIN}/${HOST_IP}
EOF
`

const unmanagedResolvConf = `
[main]
rc-manager=unmanaged
`

const forceDnsDispatcherScript = `#!/bin/bash

# In order to override cluster domain please provide this file with the following params:
# SNO_CLUSTER_NAME_OVERRIDE=<new cluster name>
# SNO_BASE_DOMAIN_OVERRIDE=<your new base domain>
# SNO_DNSMASQ_IP_OVERRIDE=<new ip>
if [ -f /etc/default/sno_dnsmasq_configuration_overrides ]; then
    source /etc/default/sno_dnsmasq_configuration_overrides
fi

HOST_IP=${SNO_DNSMASQ_IP_OVERRIDE:-"{{.HOST_IP}}"}
CLUSTER_NAME=${SNO_CLUSTER_NAME_OVERRIDE:-"{{.CLUSTER_NAME}}"}
BASE_DOMAIN=${SNO_BASE_DOMAIN_OVERRIDE:-"{{.DNS_DOMAIN}}"}
CLUSTER_FULL_DOMAIN="${CLUSTER_NAME}.${BASE_DOMAIN}"

export BASE_RESOLV_CONF=/run/NetworkManager/resolv.conf
if [ "$2" = "dhcp4-change" ] || [ "$2" = "dhcp6-change" ] || [ "$2" = "up" ] || [ "$2" = "connectivity-change" ]; then
	export TMP_FILE=$(mktemp /etc/forcedns_resolv.conf.XXXXXX)
	cp  $BASE_RESOLV_CONF $TMP_FILE
	chmod --reference=$BASE_RESOLV_CONF $TMP_FILE
	sed -i -e "s/${CLUSTER_FULL_DOMAIN}//" \
	-e "s/search /& ${CLUSTER_FULL_DOMAIN} /" \
	-e "0,/nameserver/s/nameserver/& $HOST_IP\n&/" $TMP_FILE
	mv $TMP_FILE /etc/resolv.conf
fi
`

const dnsMachineConfigManifest = `
apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  labels:
    machineconfiguration.openshift.io/role: master
  name: 50-master-dnsmasq-configuration
spec:
  config:
    ignition:
      version: 3.1.0
    storage:
      files:
        - contents:
            source: data:text/plain;charset=utf-8;base64,{{.DNSMASQ_CONTENT}}
          mode: 365
          path: /usr/local/bin/dnsmasq_config.sh
          overwrite: true
        - contents:
            source: data:text/plain;charset=utf-8;base64,{{.FORCE_DNS_SCRIPT}}
          mode: 365
          path: /etc/NetworkManager/dispatcher.d/forcedns
          overwrite: true
        - contents:
            source: data:text/plain;charset=utf-8;base64,{{.UNMANAGED_RESOLV_CONF}}
          mode: 420
          path: /etc/NetworkManager/conf.d/single-node.conf
          overwrite: true
    systemd:
      units:
        - name: dnsmasq.service
          enabled: true
          contents: |
            [Unit]
            Description=Run dnsmasq to provide local dns for Single Node OpenShift
            Before=kubelet.service crio.service
            After=network.target ovs-configuration.service

            [Service]
            TimeoutStartSec=30
            ExecStartPre=/usr/local/bin/dnsmasq_config.sh
            ExecStart=/usr/sbin/dnsmasq -k
            Restart=always

            [Install]
            WantedBy=multi-user.target
`

const schedulableMastersManifestPatch = `---
- op: replace
  path: /spec/mastersSchedulable
  value: true
`

func (m *ManifestsGenerator) createChronyManifestContent(c *common.Cluster, role models.HostRole, log logrus.FieldLogger) ([]byte, error) {
	sources := make([]string, 0)

	for _, host := range c.Hosts {
		if host.NtpSources != "" {
			var ntpSources []*models.NtpSource
			if err := json.Unmarshal([]byte(host.NtpSources), &ntpSources); err != nil {
				return nil, errors.Wrapf(err, "Failed to unmarshal %s", host.NtpSources)
			}

			for _, source := range ntpSources {
				if !lo.Contains(sources, source.SourceName) {
					sources = append(sources, source.SourceName)
				}
			}
		}

		// Avoiding a ZTP race that installation may start before getting ntp reply from the agent
		rawSources, err := common.GetHostNTPSources(m.DB, host)
		if err != nil {
			return nil, err
		}
		additionalNTPSources := strings.Split(rawSources, ",")
		for _, source := range additionalNTPSources {
			if source == "" {
				continue
			}
			if !lo.Contains(sources, source) {
				sources = append(sources, source)
			}
		}
	}

	content := defaultChronyConf[:]

	for _, source := range sources {
		content += fmt.Sprintf("\nserver %s iburst", source)
	}

	var manifestParams = map[string]interface{}{
		"CHRONY_CONTENT": base64.StdEncoding.EncodeToString([]byte(content)),
		"ROLE":           string(role),
	}

	return fillTemplate(manifestParams, ntpMachineConfigManifest, log)
}

func (m *ManifestsGenerator) AddChronyManifest(ctx context.Context, log logrus.FieldLogger, cluster *common.Cluster) error {
	roles := []models.HostRole{models.HostRoleMaster, models.HostRoleWorker}
	if common.IsClusterTopologyHighlyAvailableArbiter(cluster) {
		roles = append(roles, models.HostRoleArbiter)
	}
	for _, role := range roles {
		content, err := m.createChronyManifestContent(cluster, role, log)

		if err != nil {
			return errors.Wrapf(err, "Failed to create chrony manifest content for role %s cluster id %s", role, *cluster.ID)
		}

		chronyManifestFileName := fmt.Sprintf("50-%ss-chrony-configuration.yaml", string(role))
		err = m.createManifests(ctx, cluster, chronyManifestFileName, content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ManifestsGenerator) AddSchedulableMastersManifest(ctx context.Context, log logrus.FieldLogger, cluster *common.Cluster) error {
	content := []byte(schedulableMastersManifestPatch)
	schedulableMastersManifestFile := "cluster-scheduler-02-config.yml.patch_ai_set_masters_schedulable"
	err := m.createManifests(ctx, cluster, schedulableMastersManifestFile, content)
	if err != nil {
		return err
	}
	return nil
}

const diskEncryptionManifest = `apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  name: {{ .ROLE }}-{{ .MODE }}
  labels:
    machineconfiguration.openshift.io/role: {{ .ROLE }}
spec:
  config:
    ignition:
      version: 3.2.0
    storage:
      luks:
        - name: root
          device: /dev/disk/by-partlabel/root
          clevis:
		  {{- if eq .MODE "tpm" }}
            tpm2: true
		  {{- else if eq .MODE "tang" }}
            tang:
            {{- range .TANG_SERVERS }}
              - url: {{ .Url }}
                thumbprint: {{ .Thumbprint }}
            {{- end }}
		  {{- end }}
          options: [--cipher, aes-cbc-essiv:sha256]
          wipeVolume: true
      filesystems:
        - device: /dev/mapper/root
          format: xfs
          wipeFilesystem: true
          label: root
{{- if eq .MODE "tang" }}
  kernelArguments:
    - rd.neednet=1
{{- end }}`

func (m *ManifestsGenerator) createDiskEncryptionManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster,
	manifestParams map[string]interface{}) error {

	log.Infof("Creating manifest to encrypt installation disk on %s nodes using %s encryption", manifestParams["ROLE"], manifestParams["ROLE"])

	content, err := fillTemplate(manifestParams, diskEncryptionManifest, log)
	if err != nil {
		log.WithError(err).Error("Failed to parse disk encryption manifest's template")
		return err
	}

	filename := fmt.Sprintf("99-openshift-%s-%s-encryption.yaml", manifestParams["ROLE"], manifestParams["MODE"])
	if err := m.createManifests(ctx, c, filename, content); err != nil {
		log.WithError(err).Errorf("Failed to create manifest to encrypt installation disk")
		return err
	}

	return nil
}

func (m *ManifestsGenerator) AddDiskEncryptionManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {

	if swag.StringValue(c.DiskEncryption.EnableOn) == models.DiskEncryptionEnableOnNone {
		return nil
	}

	manifestParams := map[string]interface{}{}

	switch *c.DiskEncryption.Mode {

	case models.DiskEncryptionModeTpmv2:

		manifestParams["MODE"] = "tpm"

	case models.DiskEncryptionModeTang:

		tangServers, err := tang.UnmarshalTangServers(c.DiskEncryption.TangServers)
		if err != nil {
			log.WithError(err).Error("failed to unmarshal tang_server from cluster object")
			return err
		}

		manifestParams["MODE"] = "tang"
		manifestParams["TANG_SERVERS"] = tangServers
	}

	enabledGroups := strings.Split(swag.StringValue(c.DiskEncryption.EnableOn), ",")
	isDiskEncryptionOnAll := swag.StringValue(c.DiskEncryption.EnableOn) == models.DiskEncryptionEnableOnAll

	if isDiskEncryptionOnAll || funk.ContainsString(enabledGroups, models.DiskEncryptionEnableOnMasters) {
		manifestParams["ROLE"] = "master"
		if err := m.createDiskEncryptionManifest(ctx, log, c, manifestParams); err != nil {
			return err
		}
	}

	if (isDiskEncryptionOnAll || funk.ContainsString(enabledGroups, models.DiskEncryptionEnableOnArbiters)) &&
		common.IsClusterTopologyHighlyAvailableArbiter(c) {
		manifestParams["ROLE"] = "arbiter"
		if err := m.createDiskEncryptionManifest(ctx, log, c, manifestParams); err != nil {
			return err
		}
	}

	if isDiskEncryptionOnAll || funk.ContainsString(enabledGroups, models.DiskEncryptionEnableOnWorkers) {
		manifestParams["ROLE"] = "worker"
		if err := m.createDiskEncryptionManifest(ctx, log, c, manifestParams); err != nil {
			return err
		}
	}

	return nil
}

func (m *ManifestsGenerator) createManifests(ctx context.Context, cluster *common.Cluster, filename string, content []byte) error {
	// all relevant logs of creating manifest will be inside CreateClusterManifest
	// Mark internally created manifests as "non custom", i.e. not uploaded by a user.
	_, err := m.manifestsApi.CreateClusterManifestInternal(ctx, operations.V2CreateClusterManifestParams{
		ClusterID: *cluster.ID,
		CreateManifestParams: &models.CreateManifestParams{
			Content:  swag.String(base64.StdEncoding.EncodeToString(content)),
			FileName: &filename,
			Folder:   swag.String(models.ManifestFolderOpenshift),
		},
	}, false)

	if err != nil {
		return errors.Wrapf(err, "Failed to create manifest %s", filename)
	}

	return nil
}

func (m *ManifestsGenerator) IsSNODNSMasqEnabled() bool {
	return m.Config.EnableSingleNodeDnsmasq
}

func (m *ManifestsGenerator) AddDnsmasqForSingleNode(ctx context.Context, log logrus.FieldLogger, cluster *common.Cluster) error {
	if !m.IsSNODNSMasqEnabled() {
		return nil
	}

	filename := "dnsmasq-bootstrap-in-place.yaml"

	content, err := createDnsmasqForSingleNode(log, cluster)
	if err != nil {
		log.WithError(err).Errorf("Failed to create dnsmasq manifest")
		return err
	}

	return m.createManifests(ctx, cluster, filename, content)
}

func createDnsmasqForSingleNode(log logrus.FieldLogger, cluster *common.Cluster) ([]byte, error) {
	hostIp, err := GetIpForSingleNodeInstallation(cluster, log)
	if err != nil {
		return nil, err
	}

	var manifestParams = map[string]interface{}{
		"CLUSTER_NAME": cluster.Cluster.Name,
		"DNS_DOMAIN":   cluster.Cluster.BaseDNSDomain,
		"HOST_IP":      hostIp,
	}

	log.Infof("Creating dnsmasq manifest with values: cluster name: %q, domain - %q, host ip - %q",
		cluster.Cluster.Name, cluster.Cluster.BaseDNSDomain, hostIp)

	content, err := fillTemplate(manifestParams, snoDnsmasqConf, log)
	if err != nil {
		return nil, err
	}
	dnsmasqContent := base64.StdEncoding.EncodeToString(content)

	content, err = fillTemplate(manifestParams, forceDnsDispatcherScript, log)
	if err != nil {
		return nil, err
	}
	forceDnsDispatcherScriptContent := base64.StdEncoding.EncodeToString(content)

	manifestParams = map[string]interface{}{
		"DNSMASQ_CONTENT":       dnsmasqContent,
		"FORCE_DNS_SCRIPT":      forceDnsDispatcherScriptContent,
		"UNMANAGED_RESOLV_CONF": base64.StdEncoding.EncodeToString([]byte(unmanagedResolvConf)),
	}

	content, err = fillTemplate(manifestParams, dnsMachineConfigManifest, log)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func fillTemplate(manifestParams map[string]interface{}, templateData string, log logrus.FieldLogger) ([]byte, error) {
	tmpl, err := template.New("template").Parse(templateData)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create template")
	}
	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, manifestParams); err != nil {
		log.WithError(err).Errorf("Failed to set manifest params %v to template", manifestParams)
		return nil, err
	}
	return buf.Bytes(), nil
}

const (
	redirectTelemeterStageManifest = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: cluster-monitoring-config
  namespace: openshift-monitoring
data:
  config.yaml: |
    telemeterClient:
      telemeterServerURL: {{.TELEMETER_SERVER_URL}}
`

	stageServiceBaseURL       = "https://api.stage.openshift.com"
	integrationServiceBaseURL = "https://api.integration.openshift.com"
	stageTelemeterURL         = "https://infogw.api.stage.openshift.com"
	dummyURL                  = "https://dummy.invalid"
)

// Default Telemeter server is prod.
// In case the cluster is created in stage env we need to redirct to Telemter-stage
// Note: There is no Telemeter-integraion so in this and any other cases we will redirect the metrics to a dummy URL
func (m *ManifestsGenerator) AddTelemeterManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {

	manifestParams := map[string]interface{}{}

	switch m.Config.ServiceBaseURL {

	case stageServiceBaseURL:

		log.Infof("Creating manifest to redirect metrics from installed cluster to telemeter-stage")
		manifestParams["TELEMETER_SERVER_URL"] = stageTelemeterURL

	case integrationServiceBaseURL:

		log.Infof("Creating manifest to redirect metrics from installed cluster to dummy URL")
		manifestParams["TELEMETER_SERVER_URL"] = dummyURL

	default:
		return nil

	}

	content, err := fillTemplate(manifestParams, redirectTelemeterStageManifest, log)
	if err != nil {
		log.WithError(err).Error("Failed to parse metrics redirection's template")
		return err
	}

	if err := m.createManifests(ctx, c, "redirect-telemeter.yaml", content); err != nil {

		log.WithError(err).Error("Failed to create manifest to redirect metrics from installed cluster")
		return err
	}

	return nil
}

const nicReapplyManifest = `
apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  labels:
    machineconfiguration.openshift.io/role: {{.ROLE}}
  name: 50-{{.ROLE}}s-iscsi-nic-reapply
spec:
  config:
    ignition:
      version: 3.1.0
    systemd:
      units:
        - name: iscsi-nic-reapply.service
          enabled: true
          contents: |
            # This service is used to force the reconfiguration of the network interfaces
            # on first boot when booting from an iSCSI volume because the DNS configuration
            # is not propagated properly to the real root.
            # This is a workaround for bug https://issues.redhat.com/browse/OCPBUGS-26580.
            [Unit]
            Description=Force reapply of network configuration on first boot when iSCSI boot volume is in use
            After=NetworkManager.service

            [Service]
            Type=oneshot
            ExecStart=-/bin/sh -c ' set -x; \
            # Query to detect multipath iSCSI boot volumes \
            MULTIPATH=".blockdevices[] | select(.tran == \\\"iscsi\\\") | select(.children) | .children[] | select(.children) | .children[].mountpoints | select(.) | index(\\\"/sysroot\\\")"; \
            # Query to detect standalone iSCSI boot volumes \
            ISCSI_ALONE=".blockdevices[] | select(.tran == \\\"iscsi\\\") | select(.children) | .children[].mountpoints | select(.) | index(\\\"/sysroot\\\")"; \
            (lsblk -o NAME,TRAN,MOUNTPOINTS --json | jq -e "$MULTIPATH" > /dev/null) || (lsblk -o NAME,TRAN,MOUNTPOINTS --json | jq -e "$ISCSI_ALONE" > /dev/null); \
            if [ $? = 0 ]; \
            then \
                echo "iSCSI OR multipath iSCSI boot volume detected, force network reconfiguration..."; \
                nmcli -t -f DEVICE device status | xargs -l nmcli device reapply; \
            fi'
            ExecStartPost=-systemctl disable iscsi-nic-reapply.service

            [Install]
            WantedBy=multi-user.target
`

func (m *ManifestsGenerator) AddNicReapply(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	// Add this manifest only if one of the host is installing on an iSCSI / multiapth + iSCSI boot drive
	_, isUsingISCSIBootDrive := lo.Find(c.Cluster.Hosts, func(h *models.Host) bool {
		inventory, err := common.UnmarshalInventory(h.Inventory)
		if err != nil {
			return false
		}
		installationDisk := hostutil.GetDiskByInstallationPath(inventory.Disks, hostutil.GetHostInstallationPath(h))
		if installationDisk.DriveType == models.DriveTypeMultipath {
			iSCSIDisks := hostutil.GetDisksOfHolderByType(inventory.Disks, installationDisk, models.DriveTypeISCSI)
			return len(iSCSIDisks) > 0
		}
		return installationDisk.DriveType == models.DriveTypeISCSI
	})

	if !isUsingISCSIBootDrive {
		return nil
	}

	manifestParamsList := []map[string]interface{}{
		{"ROLE": "master"},
		{"ROLE": "worker"},
	}
	if common.IsClusterTopologyHighlyAvailableArbiter(c) {
		manifestParamsList = append(manifestParamsList, map[string]interface{}{"ROLE": "arbiter"})
	}
	for _, manifestParams := range manifestParamsList {
		content, err := fillTemplate(manifestParams, nicReapplyManifest, log)
		if err != nil {
			log.WithError(err).Error("Failed to parse nic reapply template")
			return err
		}
		manifestFilename := fmt.Sprintf("50-%ss-iscsi-nic-reapply.yaml", manifestParams["ROLE"])
		if err := m.createManifests(ctx, c, manifestFilename, content); err != nil {
			log.WithError(err).Error("Failed to create nic reqpply manifest")
			return err
		}
	}
	return nil
}

// NewConfig returns network config if env vars can be parsed
func NewConfig() (*Config, error) {
	networkCfg := Config{}
	if err := envconfig.Process("", &networkCfg); err != nil {
		return &networkCfg, errors.Wrapf(err, "failed to process env var to build network config")
	}
	return &networkCfg, nil
}
