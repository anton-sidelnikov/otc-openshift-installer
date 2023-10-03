package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/tls"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// TestWorkerGenerate tests generating the worker asset.
func TestWorkerGenerate(t *testing.T) {
	installConfig := installconfig.MakeAsset(
		&types.InstallConfig{
			Networking: &types.Networking{
				ServiceNetwork: []ipnet.IPNet{*ipnet.MustParseCIDR("10.0.1.0/24")},
			},
			Platform: types.Platform{
				OpenStack: &openstack.Platform{},
			},
		})

	rootCA := &tls.RootCA{}
	err := rootCA.Generate(nil)
	assert.NoError(t, err, "unexpected error generating root CA")

	parents := asset.Parents{}
	parents.Add(installConfig, rootCA)

	worker := &Worker{}
	err = worker.Generate(parents)
	assert.NoError(t, err, "unexpected error generating worker asset")

	actualFiles := worker.Files()
	assert.Equal(t, 1, len(actualFiles), "unexpected number of files in worker state")
	assert.Equal(t, "worker.ign", actualFiles[0].Filename, "unexpected name for worker ignition config")
}
