package types

import (
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// ClusterMetadata contains information
// regarding the cluster that was created by installer.
type ClusterMetadata struct {
	// ClusterName is the name for the cluster.
	ClusterName string `json:"clusterName"`
	// ClusterID is a globally unique ID that is used to identify an Openshift cluster.
	ClusterID string `json:"clusterID"`
	// InfraID is an ID that is used to identify cloud resources created by the installer.
	InfraID                 string `json:"infraID"`
	ClusterPlatformMetadata `json:",inline"`
}

// ClusterPlatformMetadata contains metadata for platfrom.
type ClusterPlatformMetadata struct {
	OpenStack *openstack.Metadata `json:"openstack,omitempty"`
}

// Platform returns a string representation of the platform
// (e.g. "aws" if AWS is non-nil).  It returns an empty string if no
// platform is configured.
func (cpm *ClusterPlatformMetadata) Platform() string {
	if cpm.OpenStack != nil {
		return openstack.Name
	}
	return ""
}
