package machines

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition/machine"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/rhcos"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	ostypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

func TestWorkerGenerate(t *testing.T) {
	cases := []struct {
		name                  string
		key                   string
		hyperthreading        types.HyperthreadingMode
		expectedMachineConfig []string
	}{
		{
			name:           "no key hyperthreading enabled",
			hyperthreading: types.HyperthreadingEnabled,
		},
		{
			name:           "key present hyperthreading enabled",
			key:            "ssh-rsa: dummy-key",
			hyperthreading: types.HyperthreadingEnabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-ssh
spec:
  config:
    ignition:
      version: 3.2.0
    passwd:
      users:
      - name: core
        sshAuthorizedKeys:
        - 'ssh-rsa: dummy-key'
  extensions: null
  fips: false
  kernelArguments: null
  kernelType: ""
  osImageURL: ""
`},
		},
		{
			name:           "no key hyperthreading disabled",
			hyperthreading: types.HyperthreadingDisabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-disable-hyperthreading
spec:
  config:
    ignition:
      version: 3.2.0
  extensions: null
  fips: false
  kernelArguments:
  - nosmt
  kernelType: ""
  osImageURL: ""
`},
		},
		{
			name:           "key present hyperthreading disabled",
			key:            "ssh-rsa: dummy-key",
			hyperthreading: types.HyperthreadingDisabled,
			expectedMachineConfig: []string{`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-disable-hyperthreading
spec:
  config:
    ignition:
      version: 3.2.0
  extensions: null
  fips: false
  kernelArguments:
  - nosmt
  kernelType: ""
  osImageURL: ""
`, `apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  creationTimestamp: null
  labels:
    machineconfiguration.openshift.io/role: worker
  name: 99-worker-ssh
spec:
  config:
    ignition:
      version: 3.2.0
    passwd:
      users:
      - name: core
        sshAuthorizedKeys:
        - 'ssh-rsa: dummy-key'
  extensions: null
  fips: false
  kernelArguments: null
  kernelType: ""
  osImageURL: ""
`},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			parents := asset.Parents{}
			parents.Add(
				&installconfig.ClusterID{
					UUID:    "test-uuid",
					InfraID: "test-infra-id",
				},
				installconfig.MakeAsset(
					&types.InstallConfig{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test-cluster",
						},
						SSHKey:     tc.key,
						BaseDomain: "test-domain",
						Platform: types.Platform{
							OpenStack: &ostypes.Platform{},
						},
						Compute: []types.MachinePool{
							{
								Replicas:       pointer.Int64Ptr(1),
								Hyperthreading: tc.hyperthreading,
								Platform: types.MachinePoolPlatform{
									OpenStack: &ostypes.MachinePool{
										Zones: []string{"us-east-1a"},
									},
								},
							},
						},
					}),
				(*rhcos.Image)(pointer.StringPtr("test-image")),
				(*rhcos.Release)(pointer.StringPtr("412.86.202208101040-0")),
				&machine.Worker{
					File: &asset.File{
						Filename: "worker-ignition",
						Data:     []byte("test-ignition"),
					},
				},
			)
			worker := &Worker{}
			if err := worker.Generate(parents); err != nil {
				t.Fatalf("failed to generate worker machines: %v", err)
			}
			expectedLen := len(tc.expectedMachineConfig)
			if assert.Equal(t, expectedLen, len(worker.MachineConfigFiles)) {
				for i := 0; i < expectedLen; i++ {
					assert.Equal(t, tc.expectedMachineConfig[i], string(worker.MachineConfigFiles[i].Data), "unexepcted machine config contents")
				}
			} else {
				assert.Equal(t, 0, len(worker.MachineConfigFiles), "expected no machine config files")
			}
		})
	}
}

func TestComputeIsNotModified(t *testing.T) {
	parents := asset.Parents{}
	installConfig := installconfig.MakeAsset(
		&types.InstallConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-cluster",
			},
			SSHKey:     "ssh-rsa: dummy-key",
			BaseDomain: "test-domain",
			Platform: types.Platform{
				OpenStack: &ostypes.Platform{
					DefaultMachinePlatform: &ostypes.MachinePool{},
				},
			},
			Compute: []types.MachinePool{
				{
					Replicas:       pointer.Int64Ptr(1),
					Hyperthreading: types.HyperthreadingDisabled,
					Platform: types.MachinePoolPlatform{
						OpenStack: &ostypes.MachinePool{
							Zones: []string{"us-east-1a"},
						},
					},
				},
			},
		})

	parents.Add(
		&installconfig.ClusterID{
			UUID:    "test-uuid",
			InfraID: "test-infra-id",
		},
		installConfig,
		(*rhcos.Image)(pointer.StringPtr("test-image")),
		(*rhcos.Release)(pointer.StringPtr("412.86.202208101040-0")),
		&machine.Worker{
			File: &asset.File{
				Filename: "worker-ignition",
				Data:     []byte("test-ignition"),
			},
		},
	)
	worker := &Worker{}
	if err := worker.Generate(parents); err != nil {
		t.Fatalf("failed to generate master machines: %v", err)
	}

	if installConfig.Config.Compute[0].Platform.OpenStack.FlavorName != "" {
		t.Fatalf("compute in the install config has been modified")
	}
}

func TestDefaultAWSMachinePoolPlatform(t *testing.T) {
	type testCase struct {
		name                string
		poolName            string
		expectedMachinePool ostypes.MachinePool
		assert              func(tc *testCase)
	}
	cases := []testCase{
		{
			name:     "default EBS type for compute pool",
			poolName: types.MachinePoolComputeRoleName,
			expectedMachinePool: ostypes.MachinePool{
				RootVolume: &ostypes.RootVolume{
					Size: decimalRootVolumeSize,
				},
			},
			assert: func(tc *testCase) {
				mp := defaultOpenStackMachinePoolPlatform()
				want := tc.expectedMachinePool.RootVolume.Size
				got := mp.RootVolume.Size
				assert.Equal(t, want, got, "unexpected EBS Size")
			},
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.assert(&tc)
		})
	}
}
