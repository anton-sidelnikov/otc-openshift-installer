package openstack

import (
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/destroy/providers"
)

func init() {
	providers.Registry["openstack"] = New
}
