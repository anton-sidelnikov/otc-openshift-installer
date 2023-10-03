package manifests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	configv1 "github.com/openshift/api/config/v1"
)

// installConfigFromTopologies generates an install config that would yield the
// given topologies when determineTopologies is called on it
func installConfigFromTopologies(t *testing.T, options []icOption,
	controlPlaneTopology configv1.TopologyMode, infrastructureTopology configv1.TopologyMode) *types.InstallConfig {
	installConfig := icBuild.build(options...)

	highlyAvailable := int64(3)
	singleReplica := int64(1)

	switch controlPlaneTopology {
	case configv1.HighlyAvailableTopologyMode:
		installConfig.ControlPlane = &types.MachinePool{
			Replicas: &highlyAvailable,
		}
	case configv1.SingleReplicaTopologyMode:
		installConfig.ControlPlane = &types.MachinePool{
			Replicas: &singleReplica,
		}
	}

	switch infrastructureTopology {
	case configv1.HighlyAvailableTopologyMode:
		installConfig.Compute = []types.MachinePool{
			{Replicas: &highlyAvailable},
		}
	case configv1.SingleReplicaTopologyMode:
		installConfig.Compute = []types.MachinePool{
			{Replicas: &singleReplica},
		}
	}

	// Assert that this function actually works
	generatedControlPlaneTopology, generatedInfrastructureTopology := determineTopologies(installConfig)
	assert.Equal(t, generatedControlPlaneTopology, controlPlaneTopology)
	assert.Equal(t, generatedInfrastructureTopology, infrastructureTopology)

	return installConfig
}

func TestGenerateIngerssDefaultPlacement(t *testing.T) {
	cases := []struct {
		name                        string
		installConfigBuildOptions   []icOption
		controlPlaneTopology        configv1.TopologyMode
		infrastructureTopology      configv1.TopologyMode
		expectedIngressPlacement    configv1.DefaultPlacement
		expectedIngressAWSLBType    configv1.AWSLBType
		expectedIngressPlatformType configv1.PlatformType
	}{
		{
			name:                        "test setting of openstack",
			expectedIngressPlatformType: configv1.OpenStackPlatformType,
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
					installConfigFromTopologies(t, tc.installConfigBuildOptions,
						tc.controlPlaneTopology, tc.infrastructureTopology),
				),
			)
			ingressAsset := &Ingress{}
			err := ingressAsset.Generate(parents)
			if !assert.NoError(t, err, "failed to generate asset") {
				return
			}
			if !assert.Len(t, ingressAsset.FileList, 1, "expected only one file to be generated") {
				return
			}
			assert.Equal(t, ingressAsset.FileList[0].Filename, "manifests/cluster-ingress-02-config.yml")
			var actualIngress configv1.Ingress
			err = yaml.Unmarshal(ingressAsset.FileList[0].Data, &actualIngress)
			if !assert.NoError(t, err, "failed to unmarshal infra manifest") {
				return
			}
			assert.Equal(t, tc.expectedIngressPlacement, actualIngress.Status.DefaultPlacement)
			if len(tc.expectedIngressPlatformType) != 0 {
				assert.Equal(t, tc.expectedIngressAWSLBType, actualIngress.Spec.LoadBalancer.Platform.AWS.Type)
				assert.Equal(t, tc.expectedIngressPlatformType, actualIngress.Spec.LoadBalancer.Platform.Type)
			}
		})
	}
}
