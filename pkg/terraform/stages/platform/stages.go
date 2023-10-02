package platform

import (
	"fmt"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/terraform"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/terraform/stages/alibabacloud"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/terraform/stages/openstack"
	alibabacloudtypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/alibabacloud"
	openstacktypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// StagesForPlatform returns the terraform stages to run to provision the infrastructure for the specified platform.
func StagesForPlatform(platform string) []terraform.Stage {
	switch platform {
	case alibabacloudtypes.Name:
		return alibabacloud.PlatformStages
	case openstacktypes.Name:
		return openstack.PlatformStages
	default:
		panic(fmt.Sprintf("unsupported platform %q", platform))
	}
}
