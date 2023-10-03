package platform

import (
	"fmt"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/terraform"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/terraform/stages/openstack"
	openstacktypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// StagesForPlatform returns terraform stages to run to provision the infrastructure for the specified platform.
func StagesForPlatform(platform string) []terraform.Stage {
	switch platform {
	case openstacktypes.Name:
		return openstack.PlatformStages
	default:
		panic(fmt.Sprintf("unsupported platform %q", platform))
	}
}
