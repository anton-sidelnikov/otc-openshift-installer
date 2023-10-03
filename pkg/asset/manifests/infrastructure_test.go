package manifests

import (
	ostypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	configv1 "github.com/openshift/api/config/v1"
)

func TestGenerateInfrastructure(t *testing.T) {
	cases := []struct {
		name                   string
		installConfig          *types.InstallConfig
		expectedInfrastructure *configv1.Infrastructure
	}{{
		name:          "vanilla aws",
		installConfig: icBuild.build(icBuild.forOpenstack()),
		expectedInfrastructure: infraBuild.build(
			infraBuild.forPlatform(configv1.OpenStackPlatformType),
		),
	}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			parents := asset.Parents{}
			parents.Add(
				&installconfig.ClusterID{
					UUID:    "test-uuid",
					InfraID: "test-infra-id",
				},
				installconfig.MakeAsset(tc.installConfig),
				&CloudProviderConfig{},
				&AdditionalTrustBundleConfig{},
			)
			infraAsset := &Infrastructure{}
			err := infraAsset.Generate(parents)
			if !assert.NoError(t, err, "failed to generate asset") {
				return
			}
			if !assert.Len(t, infraAsset.FileList, 1, "expected only one file to be generated") {
				return
			}
			assert.Equal(t, infraAsset.FileList[0].Filename, "manifests/cluster-infrastructure-02-config.yml")
			var actualInfra configv1.Infrastructure
			err = yaml.Unmarshal(infraAsset.FileList[0].Data, &actualInfra)
			if !assert.NoError(t, err, "failed to unmarshal infra manifest") {
				return
			}
			assert.Equal(t, tc.expectedInfrastructure, &actualInfra)
		})
	}
}

type icOption func(*types.InstallConfig)

type icBuildNamespace struct{}

var icBuild icBuildNamespace

func (icBuildNamespace) build(opts ...icOption) *types.InstallConfig {
	ic := &types.InstallConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cluster",
		},
		BaseDomain:   "test-domain",
		ControlPlane: &types.MachinePool{},
	}
	for _, opt := range opts {
		opt(ic)
	}
	return ic
}

func (b icBuildNamespace) forOpenstack() icOption {
	return func(ic *types.InstallConfig) {
		if ic.Platform.OpenStack != nil {
			return
		}
		ic.Platform.OpenStack = &ostypes.Platform{}
	}
}

func (b icBuildNamespace) withLBType(lb configv1.OpenStackPlatformLoadBalancer) icOption {
	return func(ic *types.InstallConfig) {
		b.forOpenstack()(ic)
		ic.Platform.OpenStack.LoadBalancer.Type = lb.Type
	}
}

type infraOption func(*configv1.Infrastructure)

type infraBuildNamespace struct{}

var infraBuild infraBuildNamespace

func (b infraBuildNamespace) build(opts ...infraOption) *configv1.Infrastructure {
	infra := &configv1.Infrastructure{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.SchemeGroupVersion.String(),
			Kind:       "Infrastructure",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
		},
		Spec: configv1.InfrastructureSpec{
			PlatformSpec: configv1.PlatformSpec{},
		},
		Status: configv1.InfrastructureStatus{
			InfrastructureName:     "test-infra-id",
			APIServerURL:           "https://api.test-cluster.test-domain:6443",
			APIServerInternalURL:   "https://api-int.test-cluster.test-domain:6443",
			ControlPlaneTopology:   configv1.HighlyAvailableTopologyMode,
			InfrastructureTopology: configv1.HighlyAvailableTopologyMode,
			PlatformStatus:         &configv1.PlatformStatus{},
			CPUPartitioning:        configv1.CPUPartitioningNone,
		},
	}
	for _, opt := range opts {
		opt(infra)
	}
	return infra
}

func (b infraBuildNamespace) forPlatform(platform configv1.PlatformType) infraOption {
	return func(infra *configv1.Infrastructure) {
		infra.Spec.PlatformSpec.Type = platform
		infra.Status.PlatformStatus.Type = platform
		infra.Status.Platform = platform
	}
}
