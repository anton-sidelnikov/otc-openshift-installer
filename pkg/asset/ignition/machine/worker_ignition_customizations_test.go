package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/tls"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/ipnet"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// TestWorkerIgnitionCustomizationsGenerate tests generating the worker ignition check asset.
func TestWorkerIgnitionCustomizationsGenerate(t *testing.T) {
	cases := []struct {
		name          string
		customize     bool
		assetExpected bool
	}{
		{
			name:          "not customized",
			customize:     false,
			assetExpected: false,
		},
		{
			name:          "pointer customized",
			customize:     true,
			assetExpected: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
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

			if tc.customize == true {
				// Modify the worker config, emulating a customization to the pointer
				worker.Config.Storage.Files = append(worker.Config.Storage.Files,
					ignition.FileFromString("/etc/foo", "root", 0644, "foo"))
			}

			parents.Add(worker)
			workerIgnCheck := &WorkerIgnitionCustomizations{}
			err = workerIgnCheck.Generate(parents)
			assert.NoError(t, err, "unexpected error generating worker ignition check asset")

			actualFiles := workerIgnCheck.Files()
			if tc.assetExpected == true {
				assert.Equal(t, 1, len(actualFiles), "unexpected number of files in worker state")
				assert.Equal(t, workerMachineConfigFileName, actualFiles[0].Filename, "unexpected name for worker ignition config")
			} else {
				assert.Equal(t, 0, len(actualFiles), "unexpected number of files in worker state")
			}
		})
	}
}
