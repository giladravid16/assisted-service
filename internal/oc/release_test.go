package oc

import (
	_ "embed"
	"fmt"
	os "os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/system"
	"github.com/openshift/assisted-service/pkg/executer"
	"github.com/openshift/assisted-service/pkg/mirrorregistries"
	logrus "github.com/sirupsen/logrus"
)

var (
	log                    = logrus.New()
	releaseImage           = "release_image"
	releaseImageMirror     = "release_image_mirror"
	cacheDir               = "/tmp"
	pullSecret             = "pull secret"
	fullVersion            = "4.6.0-0.nightly-2020-08-31-220837"
	mcoImage               = "mco_image"
	mustGatherImage        = "must_gather_image"
	baremetalInstallBinary = "openshift-baremetal-install"
)

//go:embed test_skopeo_multiarch_image_output
var test_skopeo_multiarch_image_output string

var _ = Describe("oc", func() {
	var (
		oc             Release
		tempFilePath   string
		ctrl           *gomock.Controller
		mockExecuter   *executer.MockExecuter
		mockSystemInfo *system.MockSystemInfo
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockExecuter = executer.NewMockExecuter(ctrl)
		mockSystemInfo = system.NewMockSystemInfo(ctrl)
		config := Config{MaxTries: DefaultTries, RetryDelay: time.Millisecond}
		mirrorRegistriesBuilder := mirrorregistries.New(false)
		oc = NewRelease(mockExecuter, config, mirrorRegistriesBuilder, mockSystemInfo)
		tempFilePath = "/tmp/pull-secret"
		mockExecuter.EXPECT().TempFile(gomock.Any(), gomock.Any()).DoAndReturn(
			func(dir, pattern string) (*os.File, error) {
				tempPullSecretFile, err := os.Create(tempFilePath)
				Expect(err).ShouldNot(HaveOccurred())
				return tempPullSecretFile, nil
			},
		).AnyTimes()
	})

	AfterEach(func() {
		os.Remove(tempFilePath)
	})

	Context("GetMCOImage", func() {
		It("mco image from release image", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mcoImageName, false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mcoImage, "", 0).Times(1)

			mco, err := oc.GetMCOImage(log, releaseImage, "", pullSecret)
			Expect(mco).Should(Equal(mcoImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("mco image from release image mirror", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mcoImageName, true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mcoImage, "", 0).Times(1)

			mco, err := oc.GetMCOImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mco).Should(Equal(mcoImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("mco image exists in cache", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mcoImageName, true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mcoImage, "", 0).Times(1)

			mco, err := oc.GetMCOImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mco).Should(Equal(mcoImage))
			Expect(err).ShouldNot(HaveOccurred())

			// Fetch image again
			mco, err = oc.GetMCOImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mco).Should(Equal(mcoImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("mco image with no release image or mirror", func() {
			mco, err := oc.GetMCOImage(log, "", "", pullSecret)
			Expect(mco).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})

		It("stdout with new lines", func() {
			stdout := fmt.Sprintf("\n%s\n", mcoImage)

			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mcoImageName, false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(stdout, "", 0).Times(1)

			mco, err := oc.GetMCOImage(log, releaseImage, "", pullSecret)
			Expect(mco).Should(Equal(mcoImage))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("GetMustGatherImage", func() {
		It("must-gather image from release image", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mustGatherImageName, false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mustGatherImage, "", 0).Times(1)

			mustGather, err := oc.GetMustGatherImage(log, releaseImage, "", pullSecret)
			Expect(mustGather).Should(Equal(mustGatherImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("must-gather image from release image mirror", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mustGatherImageName, true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mustGatherImage, "", 0).Times(1)

			mustGather, err := oc.GetMustGatherImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mustGather).Should(Equal(mustGatherImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("must-gather image exists in cache", func() {
			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mustGatherImageName, true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(mustGatherImage, "", 0).Times(1)

			mustGather, err := oc.GetMustGatherImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mustGather).Should(Equal(mustGatherImage))
			Expect(err).ShouldNot(HaveOccurred())

			// Fetch image again
			mustGather, err = oc.GetMustGatherImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(mustGather).Should(Equal(mustGatherImage))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("must-gather image with no release image or mirror", func() {
			mustGather, err := oc.GetMustGatherImage(log, "", "", pullSecret)
			Expect(mustGather).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})

		It("stdout with new lines", func() {
			stdout := fmt.Sprintf("\n%s\n", mustGatherImage)

			command := fmt.Sprintf(templateGetImage+" --registry-config=%s",
				mustGatherImageName, false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(stdout, "", 0).Times(1)

			mustGather, err := oc.GetMustGatherImage(log, releaseImage, "", pullSecret)
			Expect(mustGather).Should(Equal(mustGatherImage))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("GetOpenshiftVersion", func() {
		It("image version from release image", func() {
			command := fmt.Sprintf(templateGetVersion+" --registry-config=%s",
				false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(fullVersion, "", 0).Times(1)

			version, err := oc.GetOpenshiftVersion(log, releaseImage, "", pullSecret)
			Expect(version).Should(Equal(fullVersion))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("image version from release image mirror", func() {
			command := fmt.Sprintf(templateGetVersion+" --registry-config=%s",
				true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(fullVersion, "", 0).Times(1)

			version, err := oc.GetOpenshiftVersion(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(version).Should(Equal(fullVersion))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("image version with no release image or mirror", func() {
			version, err := oc.GetOpenshiftVersion(log, "", "", pullSecret)
			Expect(version).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("GetMajorMinorVersion", func() {
		tests := []struct {
			fullVersion  string
			shortVersion string
			isValid      bool
		}{
			{
				fullVersion:  "4.6.0",
				shortVersion: "4.6",
				isValid:      true,
			},
			{
				fullVersion:  "4.6.4",
				shortVersion: "4.6",
				isValid:      true,
			},
			{
				fullVersion:  "4.6",
				shortVersion: "4.6",
				isValid:      true,
			},
			{
				fullVersion:  "4.6.0-0.nightly-2020-08-31-220837",
				shortVersion: "4.6",
				isValid:      true,
			},
			{
				fullVersion: "-44",
				isValid:     false,
			},
		}

		for i := range tests {
			t := tests[i]
			It(t.fullVersion, func() {
				command := fmt.Sprintf(templateGetVersion+" --registry-config=%s",
					false, releaseImage, "", tempFilePath)
				args := splitStringToInterfacesArray(command)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(t.fullVersion, "", 0).Times(1)

				version, err := oc.GetMajorMinorVersion(log, releaseImage, "", pullSecret)

				if t.isValid {
					Expect(err).ShouldNot(HaveOccurred())
					Expect(version).Should(Equal(t.shortVersion))
				} else {
					Expect(err).Should(HaveOccurred())
					Expect(version).Should(BeEmpty())
				}
			})
		}
	})

	Context("GetReleaseArchitecture", func() {
		Context("for single-arch release image", func() {
			It("fetch cpu architecture", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				args := splitStringToInterfacesArray(command)
				imageInfoStr := fmt.Sprintf("{ \"config\": { \"architecture\": \"%s\" }}", common.TestDefaultConfig.CPUArchitecture)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(imageInfoStr, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(Equal([]string{common.TestDefaultConfig.CPUArchitecture}))
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("fail with malformed cpu architecture", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				args := splitStringToInterfacesArray(command)
				imageInfoStr := fmt.Sprintf("{ \"config\": { \"not-an-architecture\": \"%s\" }}", common.TestDefaultConfig.CPUArchitecture)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return(imageInfoStr, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(BeEmpty())
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("for multi-arch release image", func() {
			It("fetch cpu architecture", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				command2 := fmt.Sprintf(templateSkopeoDetectMultiarch+" --authfile %s", releaseImage, tempFilePath)
				args := splitStringToInterfacesArray(command)
				args2 := splitStringToInterfacesArray(command2)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "the image is a manifest list", 1).Times(1)
				mockExecuter.EXPECT().Execute(args2[0], args2[1:]...).Return(test_skopeo_multiarch_image_output, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(ConsistOf([]string{"x86_64", "ppc64le", "s390x", "arm64"}))
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("fail with malformed manifests - not a list", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				command2 := fmt.Sprintf(templateSkopeoDetectMultiarch+" --authfile %s", releaseImage, tempFilePath)
				args := splitStringToInterfacesArray(command)
				args2 := splitStringToInterfacesArray(command2)
				imageInfoStr := fmt.Sprintf("{ \"manifests\": { \"platform\": { \"not-an-architecture\": \"%s\" }}}", common.TestDefaultConfig.CPUArchitecture)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "the image is a manifest list", 1).Times(1)
				mockExecuter.EXPECT().Execute(args2[0], args2[1:]...).Return(imageInfoStr, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(BeEmpty())
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("failed to get image info using oc"))
			})

			It("fail with malformed manifests - no architecture", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				command2 := fmt.Sprintf(templateSkopeoDetectMultiarch+" --authfile %s", releaseImage, tempFilePath)
				args := splitStringToInterfacesArray(command)
				args2 := splitStringToInterfacesArray(command2)
				imageInfoStr := fmt.Sprintf("{ \"manifests\": [{ \"platform\": { \"not-an-architecture\": \"%s\" }}]}", common.TestDefaultConfig.CPUArchitecture)
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "the image is a manifest list", 1).Times(1)
				mockExecuter.EXPECT().Execute(args2[0], args2[1:]...).Return(imageInfoStr, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(BeEmpty())
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("image manifest does not contain architecture"))
			})

			It("fail with malformed manifests - empty architecture", func() {
				command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
				command2 := fmt.Sprintf(templateSkopeoDetectMultiarch+" --authfile %s", releaseImage, tempFilePath)
				args := splitStringToInterfacesArray(command)
				args2 := splitStringToInterfacesArray(command2)
				imageInfoStr := "{ \"manifests\": [{ \"platform\": { \"architecture\": \"\" }}]}"
				mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "the image is a manifest list", 1).Times(1)
				mockExecuter.EXPECT().Execute(args2[0], args2[1:]...).Return(imageInfoStr, "", 0).Times(1)

				arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
				Expect(arch).Should(BeEmpty())
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("image manifest does not contain architecture"))
			})
		})

		It("broken release image", func() {
			command := fmt.Sprintf(templateImageInfo+" --registry-config=%s", releaseImage, "", tempFilePath)
			command2 := fmt.Sprintf(templateSkopeoDetectMultiarch+" --authfile %s", releaseImage, tempFilePath)
			args := splitStringToInterfacesArray(command)
			args2 := splitStringToInterfacesArray(command2)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "that's not even an image", 1).Times(1)
			mockExecuter.EXPECT().Execute(args2[0], args2[1:]...).Return("", "that's still not an image", 1).Times(1)

			arch, err := oc.GetReleaseArchitecture(log, releaseImage, "", pullSecret)
			Expect(arch).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})

		It("no release image", func() {
			arch, err := oc.GetReleaseArchitecture(log, "", "", pullSecret)
			Expect(arch).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Extract", func() {
		BeforeEach(func() {
			mockSystemInfo.EXPECT().FIPSEnabled().Return(false, nil).AnyTimes()
		})

		It("extract baremetal-install from release image", func() {
			command := fmt.Sprintf(templateExtract+" --registry-config=%s",
				baremetalInstallBinary, filepath.Join(cacheDir, releaseImage), false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "", 0).Times(1)

			path, err := oc.Extract(log, releaseImage, "", cacheDir, pullSecret, "4.15.0")
			filePath := filepath.Join(cacheDir+"/"+releaseImage, baremetalInstallBinary)
			Expect(path).To(Equal(filePath))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("extract baremetal-install from release image mirror", func() {
			command := fmt.Sprintf(templateExtract+" --registry-config=%s",
				baremetalInstallBinary, filepath.Join(cacheDir, releaseImage), true, releaseImageMirror, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "", 0).Times(1)

			path, err := oc.Extract(log, releaseImage, releaseImageMirror, cacheDir, pullSecret, "4.15.0")
			filePath := filepath.Join(cacheDir+"/"+releaseImage, baremetalInstallBinary)
			Expect(path).To(Equal(filePath))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("extract baremetal-install with no release image or mirror", func() {
			path, err := oc.Extract(log, "", "", cacheDir, pullSecret, "4.15.0")
			Expect(path).Should(BeEmpty())
			Expect(err).Should(HaveOccurred())
		})
		It("extract baremetal-install from release image with retry", func() {
			command := fmt.Sprintf(templateExtract+" --registry-config=%s",
				baremetalInstallBinary, filepath.Join(cacheDir, releaseImage), false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "Failed to extract the installer", 1).Times(1)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "", 0).Times(1)

			path, err := oc.Extract(log, releaseImage, "", cacheDir, pullSecret, "4.15.0")
			filePath := filepath.Join(cacheDir+"/"+releaseImage, baremetalInstallBinary)
			Expect(path).To(Equal(filePath))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("extract baremetal-install from release image retry exhausted", func() {
			command := fmt.Sprintf(templateExtract+" --registry-config=%s",
				baremetalInstallBinary, filepath.Join(cacheDir, releaseImage), false, releaseImage, "", tempFilePath)
			args := splitStringToInterfacesArray(command)
			mockExecuter.EXPECT().Execute(args[0], args[1:]...).Return("", "Failed to extract the installer", 1).Times(5)

			path, err := oc.Extract(log, releaseImage, "", cacheDir, pullSecret, "4.15.0")
			Expect(path).To(Equal(""))
			Expect(err).Should(HaveOccurred())
		})
	})
	Context("GetCoreOSImage", func() {
		It("should return rhel-coreos for OCP image", func() {
			expectedForImage := "rhel-coreos"
			image := "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3c8f6006b25b4c9aef0cfb12de5e44134317f5347bff85eba35af22ad23d964f"
			mockExecuter.EXPECT().Execute(
				"oc",
				"adm",
				"release",
				"info",
				"--image-for="+expectedForImage,
				"--insecure=true",
				releaseImageMirror,
				"--registry-config="+tempFilePath,
			).
				Return(image, "", 0).
				Times(1)
			coreosImage, err := oc.GetCoreOSImage(log, releaseImage, releaseImageMirror, pullSecret)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coreosImage).Should(Equal(image))
		})
		It("should return stream-coreos for OKD image", func() {
			streamCoreOSImage := "quay.io/okd/scos-content@sha256:6e78761cdab37d31967c91346458a7981da8a37716c815e5063bea569a3cc43d"
			mockExecuter.EXPECT().Execute(
				"oc",
				"adm",
				"release",
				"info",
				"--image-for=rhel-coreos",
				"--insecure=true",
				releaseImageMirror,
				"--registry-config="+tempFilePath,
			).
				Return("", "Error retrieving rhel-coreos image", 1).
				Times(1)
			mockExecuter.EXPECT().Execute(
				"oc",
				"adm",
				"release",
				"info",
				"--image-for=stream-coreos",
				"--insecure=true",
				releaseImageMirror,
				"--registry-config="+tempFilePath,
			).
				Times(0)
			mockExecuter.EXPECT().Execute(
				"oc",
				"adm",
				"release",
				"info",
				"--image-for=stream-coreos",
				"--insecure=true",
				releaseImageMirror,
				"--registry-config="+tempFilePath,
			).
				Return(streamCoreOSImage, "", 0).
				Times(1)
			okdReleaseImage := "quay.io/okd/scos-release:4.19.0-okd-scos.ec.6"
			coreosImage, err := oc.GetCoreOSImage(log, okdReleaseImage, releaseImageMirror, pullSecret)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coreosImage).Should(Equal(streamCoreOSImage))
		})
	})
})

var _ = Describe("getImageFromRelease", func() {
	var (
		oc                                *release
		tempFilePath                      string
		ctrl                              *gomock.Controller
		mockExecuter                      *executer.MockExecuter
		mockMirrorRegistriesConfigBuilder *mirrorregistries.MockServiceMirrorRegistriesConfigBuilder
		log                               logrus.FieldLogger
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockExecuter = executer.NewMockExecuter(ctrl)
		mockMirrorRegistriesConfigBuilder = mirrorregistries.NewMockServiceMirrorRegistriesConfigBuilder(ctrl)
		mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(false).AnyTimes()
		config := Config{MaxTries: DefaultTries, RetryDelay: time.Millisecond}
		oc = &release{
			executer:                mockExecuter,
			mirrorRegistriesBuilder: mockMirrorRegistriesConfigBuilder,
			config:                  config,
			imagesMap:               common.NewExpiringCache(time.Hour, time.Hour),
		}
		log = logrus.New()
		tempFilePath = "/tmp/pull-secret"
		mockExecuter.EXPECT().TempFile(gomock.Any(), gomock.Any()).DoAndReturn(
			func(dir, pattern string) (*os.File, error) {
				tempPullSecretFile, err := os.Create(tempFilePath)
				Expect(err).ShouldNot(HaveOccurred())
				return tempPullSecretFile, nil
			},
		).AnyTimes()
	})

	type requester struct {
		imageName      string
		releaseName    string
		expectedResult string
		timesToRun     int
	}
	tests := []struct {
		name       string
		requesters []requester
	}{
		{
			name: "Empty",
		},
		{
			name: "Single requester",
			requesters: []requester{
				{
					imageName:      "image1",
					releaseName:    "release1",
					expectedResult: "result1",
					timesToRun:     1,
				},
			},
		},
		{
			name: "Multiple requesters",
			requesters: []requester{
				{
					imageName:      "image1",
					releaseName:    "release1",
					expectedResult: "result1",
					timesToRun:     20,
				},
			},
		},
		{
			name: "Multiple requesters - two images",
			requesters: []requester{
				{
					imageName:      "image1",
					releaseName:    "release1",
					expectedResult: "result1",
					timesToRun:     20,
				},
				{
					imageName:      "image2",
					releaseName:    "release2",
					expectedResult: "result2",
					timesToRun:     20,
				},
			},
		},
	}
	for i := range tests {
		t := tests[i]
		It(t.name, func() {
			for _, r := range t.requesters {
				mockExecuter.EXPECT().Execute("oc", "adm", "release", "info",
					"--image-for="+r.imageName, "--insecure=false", r.releaseName,
					"--registry-config=/tmp/pull-secret").
					Return(r.expectedResult, "", 0).Times(1)
			}
			panicChan := make(chan interface{})
			doneChan := make(chan bool)
			numRequesting := 0
			for l := range t.requesters {
				r := t.requesters[l]
				for j := 0; j != r.timesToRun; j++ {
					numRequesting++
					go func() {
						defer func() {
							if panicVar := recover(); panicVar != nil {
								panicChan <- panicVar
							}
							doneChan <- true
						}()
						ret, err := oc.getImageFromRelease(log, r.imageName, r.releaseName, "", "pull")
						Expect(err).ToNot(HaveOccurred())
						Expect(ret).To(Equal(r.expectedResult))
					}()
				}
			}
			for numRequesting > 0 {
				select {
				case panicVar := <-panicChan:
					panic(panicVar)
				case <-doneChan:
					numRequesting--
				}
			}
		})
	}

	AfterEach(func() {
		ctrl.Finish()
	})
})

var _ = Describe("Mirrors configuration generation", func() {
	var (
		oc                                *release
		mockMirrorRegistriesConfigBuilder *mirrorregistries.MockServiceMirrorRegistriesConfigBuilder
		ctrl                              *gomock.Controller
		mockExecuter                      *executer.MockExecuter
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockExecuter = executer.NewMockExecuter(ctrl)
		mockMirrorRegistriesConfigBuilder = mirrorregistries.NewMockServiceMirrorRegistriesConfigBuilder(ctrl)
		config := Config{MaxTries: DefaultTries, RetryDelay: time.Millisecond}
		oc = &release{executer: mockExecuter, config: config, imagesMap: common.NewExpiringCache(time.Hour, time.Hour),
			mirrorRegistriesBuilder: mockMirrorRegistriesConfigBuilder}
		log = logrus.New()
	})

	DescribeTable(
		"Uses '--idms-file' if supported, for all commands",
		func(commandTemplate, helpCommand string) {
			helpArgs := whiteSpaceRE.Split(helpCommand, -1)
			mockExecuter.EXPECT().Execute(helpArgs[0], helpArgs[1:]).
				Return(`--idms-file='':`, "", 0).
				Times(1)
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, commandTemplate)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(MatchYAML(common.Dedent(`
				apiVersion: config.openshift.io/v1
				kind: ImageDigestMirrorSet
				metadata:
				  creationTimestamp: null
				  name: image-mirror-set
				spec:
				  imageDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: registry.ci.org
				status: {}
			`)))
		},
		Entry(
			"Get image",
			templateGetImage,
			"oc adm release info --help",
		),
		Entry(
			"Get version",
			templateGetVersion,
			"oc adm release info --help",
		),
		Entry(
			"Extract",
			templateExtract,
			"oc adm release extract --help",
		),
		Entry(
			"Image info",
			templateImageInfo,
			"oc image info --help",
		),
	)

	DescribeTable(
		"Uses '--icsp-file' if '--idms-file' isn't supported, for all commands",
		func(commandTemplate, helpCommand string) {
			helpArgs := whiteSpaceRE.Split(helpCommand, -1)
			mockExecuter.EXPECT().Execute(helpArgs[0], helpArgs[1:]).
				Return(`--icsp-file='':`, "", 0).
				Times(1)
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, commandTemplate)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(MatchYAML(common.Dedent(`
				apiVersion: operator.openshift.io/v1alpha1
				kind: ImageContentSourcePolicy
				metadata:
				  creationTimestamp: null
				  name: image-policy
				spec:
				  repositoryDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: registry.ci.org
			`)))
		},
		Entry(
			"Get image",
			templateGetImage,
			"oc adm release info --help",
		),
		Entry(
			"Get version",
			templateGetVersion,
			"oc adm release info --help",
		),
		Entry(
			"Extract",
			templateExtract,
			"oc adm release extract --help",
		),
		Entry(
			"Image info",
			templateImageInfo,
			"oc image info --help",
		),
	)

	Context("With IDMS support", func() {
		BeforeEach(func() {
			mockExecuter.EXPECT().Execute("oc", []string{"adm", "release", "info", "--help"}).
				Return(`--idms-file='':`, "", 0).
				Times(1)
		})

		It("One valid mirror registry", func() {
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
				{
					Location: "quay.io",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(MatchYAML(common.Dedent(`
				apiVersion: config.openshift.io/v1
				kind: ImageDigestMirrorSet
				metadata:
				  creationTimestamp: null
				  name: image-mirror-set
				spec:
				  imageDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: registry.ci.org
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: quay.io
				status: {}
			`)))
		})

		It("Multiple valid mirror registries", func() {
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
						"host1.example.org:5000/openshift",
					},
				},
				{
					Location: "quay.io",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(MatchYAML(common.Dedent(`
				apiVersion: config.openshift.io/v1
				kind: ImageDigestMirrorSet
				metadata:
				  creationTimestamp: null
				  name: image-mirror-set
				spec:
				  imageDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    - host1.example.org:5000/openshift
				    source: registry.ci.org
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: quay.io
				status: {}
			`)))
		})

		It("No registries", func() {
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(false).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(mirrorsFlag).Should(BeNil())
		})

		It("Invalid mirror registries", func() {
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(nil, fmt.Errorf("extract failed")).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("extract failed"))
			Expect(mirrorsFlag).Should(BeNil())
		})
	})

	Context("Without IDMS support, but with deprecated ICSP support", func() {
		BeforeEach(func() {
			mockExecuter.EXPECT().Execute("oc", []string{"adm", "release", "info", "--help"}).Return(
				`--icsp-file='':`, "", 0,
			).AnyTimes()
		})

		It("One valid mirror registry", func() {
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
				{
					Location: "quay.io",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(MatchYAML(common.Dedent(`
				apiVersion: operator.openshift.io/v1alpha1
				kind: ImageContentSourcePolicy
				metadata:
				  creationTimestamp: null
				  name: image-policy
				spec:
				  repositoryDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: registry.ci.org
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: quay.io
			`)))
		})

		It("Multiple valid mirror registries", func() {
			regData := []mirrorregistries.RegistriesConf{
				{
					Location: "registry.ci.org",
					Mirror: []string{
						"host1.example.org:5000/localimages",
						"host1.example.org:5000/openshift",
					},
				},
				{
					Location: "quay.io",
					Mirror: []string{
						"host1.example.org:5000/localimages",
					},
				},
			}
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(regData, nil).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			defer mirrorsFlag.Delete()
			Expect(mirrorsFlag).ShouldNot(BeNil())
			Expect(mirrorsFlag.file).ShouldNot(BeEmpty())
			data, err := os.ReadFile(mirrorsFlag.file)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(MatchYAML(common.Dedent(`
				apiVersion: operator.openshift.io/v1alpha1
				kind: ImageContentSourcePolicy
				metadata:
				  creationTimestamp: null
				  name: image-policy
				spec:
				  repositoryDigestMirrors:
				  - mirrors:
				    - host1.example.org:5000/localimages
				    - host1.example.org:5000/openshift
				    source: registry.ci.org
				  - mirrors:
				    - host1.example.org:5000/localimages
				    source: quay.io
			`)))
		})

		It("No registries", func() {
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(false).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(mirrorsFlag).Should(BeNil())
		})

		It("Invalid mirror registries", func() {
			mockMirrorRegistriesConfigBuilder.EXPECT().IsMirrorRegistriesConfigured().Return(true).Times(1)
			mockMirrorRegistriesConfigBuilder.EXPECT().ExtractLocationMirrorDataFromRegistries().Return(nil, fmt.Errorf("extract failed")).Times(1)
			mirrorsFlag, err := oc.getMirrorsFlagFromRegistriesConfig(log, templateGetImage)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("extract failed"))
			Expect(mirrorsFlag).Should(BeNil())
		})
	})
})

var _ = Describe("GetReleaseBinaryPath", func() {
	var (
		ctrl           *gomock.Controller
		mockSystemInfo *system.MockSystemInfo
		r              Release
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockSystemInfo = system.NewMockSystemInfo(ctrl)
		r = NewRelease(nil, Config{}, nil, mockSystemInfo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("with FIPS disabled", func() {
		BeforeEach(func() {
			mockSystemInfo.EXPECT().FIPSEnabled().Return(false, nil).AnyTimes()
		})

		It("returns the openshift-baremetal-install binary for versions earlier than 4.16", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.15.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
		})

		It("returns the openshift-install binary for 4.16.0", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.16.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
		})

		It("returns the openshift-install binary for 4.16 pre release", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.16.0-ec.6")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
			_, bin, _, err = r.GetReleaseBinaryPath("image", "dir", "4.16.0-rc.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
			_, bin, _, err = r.GetReleaseBinaryPath("image", "dir", "4.16.0-rc.3")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
		})

		It("returns the openshift-install binary for 4.16 nightlies", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.16.0-0.nightly-2024-05-30-130713")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
		})

		It("returns the openshift-install binary for versions later than 4.16.0", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.17.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
			_, bin, _, err = r.GetReleaseBinaryPath("image", "dir", "4.18.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-install"))
		})
	})

	Context("with FIPS enabled", func() {
		BeforeEach(func() {
			mockSystemInfo.EXPECT().FIPSEnabled().Return(true, nil).AnyTimes()
		})

		It("returns the openshift-baremetal-install binary for versions earlier than 4.16", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.15.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
		})

		It("returns the openshift-baremetal-install binary for 4.16.0", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.16.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
		})

		It("returns the openshift-install binary for 4.16 pre release", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.16.0-ec.6")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
			_, bin, _, err = r.GetReleaseBinaryPath("image", "dir", "4.16.0-rc.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
		})

		It("returns the openshift-install binary for versions later than 4.16.0", func() {
			_, bin, _, err := r.GetReleaseBinaryPath("image", "dir", "4.17.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
			_, bin, _, err = r.GetReleaseBinaryPath("image", "dir", "4.18.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bin).To(Equal("openshift-baremetal-install"))
		})
	})
})

var _ = Describe("Get flags from help", func() {
	var (
		ctrl         *gomock.Controller
		mockExecuter *executer.MockExecuter
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockExecuter = executer.NewMockExecuter(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	DescribeTable(
		"With correct help text",
		func(command string, helpCommand string, helpText string, expectedFlags []string) {
			args := strings.Split(helpCommand, " ")
			mockExecuter.EXPECT().Execute(args[0], args[1:]).Return(helpText, "", 0).AnyTimes()
			actualFlags, err := getFlagsFromHelp(log, mockExecuter, command)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualFlags).To(Equal(expectedFlags))
		},
		Entry(
			"Simple command without flags",
			"my",
			"my --help",
			`
			--my-flag='':
				Does something.
			--your-flag=true:
				Does something else.
			`,
			[]string{
				"my-flag",
				"your-flag",
			},
		),
		Entry(
			"Simple command with flags",
			"my --my-flag=my-value --your-flag=false",
			"my --help",
			`
			--my-flag='':
				Does something.
			--your-flag=true:
				Does something else.
			`,
			[]string{
				"my-flag",
				"your-flag",
			},
		),
		Entry(
			"Sub-command without flags",
			"my sub",
			"my sub --help",
			`
			--my-flag='':
				Does something.
			--your-flag=true:
				Does something else.
			`,
			[]string{
				"my-flag",
				"your-flag",
			},
		),
		Entry(
			"Sub-command with flags",
			"my sub --my-flag=my-value --your-flag=false",
			"my sub --help",
			`
			--my-flag='':
				Does something.
			--your-flag=true:
				Does something else.
			`,
			[]string{
				"my-flag",
				"your-flag",
			},
		),
		Entry(
			"Sorted results",
			"my",
			"my --help",
			`
			--flag-c='':
				Flag C.
			--flag-a=true:
				Flag A.
			--flag-b=123:
				Flag B.
			`,
			[]string{
				"flag-a",
				"flag-b",
				"flag-c",
			},
		),
		Entry(
			"Command with short option",
			"my",
			"my --help",
			`
			-a, --flag-a='':
				A command with a short option.
			`,
			[]string{
				"flag-a",
			},
		),
		Entry(
			"Ignores flag mentioned in another flag",
			"my",
			"my --help",
			`
			--my-flag='':
				Do not confuse with --your-flag.
			`,
			[]string{
				"my-flag",
			},
		),
		Entry(
			"Ignores flag starting help text",
			"my",
			"my --help",
			`
			--my-flag='':
				--your-flag='' should be ignored.
			`,
			[]string{
				"my-flag",
			},
		),
		Entry(
			"Ignores apparently correct flags inside help text",
			"my",
			"my --help",
			`
			--my-flag='':
				Ignore the --your-flag='': flag.
			`,
			[]string{
				"my-flag",
			},
		),
		Entry(
			"Respects case",
			"my",
			"my --help",
			`
			--MY-FLAG='':
				Does something.
			`,
			[]string{
				"MY-FLAG",
			},
		),
		Entry(
			"Ignores flags inside text",
			"my",
			"my --help",
			`
			--MY-FLAG='':
				Does something.
			`,
			[]string{
				"MY-FLAG",
			},
		),
	)

	It("Fails if command returns non zero exit code", func() {
		mockExecuter.EXPECT().Execute(gomock.Any(), gomock.Any()).Return("", "", 1).AnyTimes()
		flags, err := getFlagsFromHelp(log, mockExecuter, "my")
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(Equal("command 'my --help' finished with exit code 1"))
		Expect(flags).To(BeNil())
	})

	It("Fails if command doesn't exist", func() {
		// This is the contract of the executer.Executer interface: if the command doesn't exist it returns
		// -1 as the exit code and the error converted to string as the standard output.
		mockExecuter.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(
			"", `exec: "my" executalbe fine not found in $PATH`, -1,
		).AnyTimes()
		flags, err := getFlagsFromHelp(log, mockExecuter, "my")
		Expect(err).To(HaveOccurred())
		message := err.Error()
		Expect(message).To(Equal("binary of command 'my --help' doesn't exist"))
		Expect(flags).To(BeNil())
	})
})

func splitStringToInterfacesArray(str string) []interface{} {
	argsAsString := whiteSpaceRE.Split(str, -1)
	argsAsInterface := make([]interface{}, len(argsAsString))
	for i, v := range argsAsString {
		argsAsInterface[i] = v
	}

	return argsAsInterface
}

func TestOC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "oc tests")
}
