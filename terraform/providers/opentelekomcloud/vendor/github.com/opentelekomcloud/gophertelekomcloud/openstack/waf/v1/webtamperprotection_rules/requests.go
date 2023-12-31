package webtamperprotection_rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToWebTamperCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new web tamper protection rule.
type CreateOpts struct {
	Hostname string `json:"hostname" required:"true"`
	Path     string `json:"path" required:"true"`
}

// ToWebTamperCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToWebTamperCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new web tamper protection rule based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, policyID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToWebTamperCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c, policyID), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular web tamper protection rule based on its unique ID.
func Get(c *golangsdk.ServiceClient, policyID, ruleID string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, policyID, ruleID), &r.Body, openstack.StdRequestOpts())
	return
}

// Delete will permanently delete a particular web tamper protection rule based on its unique ID.
func Delete(c *golangsdk.ServiceClient, policyID, ruleID string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders}
	_, r.Err = c.Delete(resourceURL(c, policyID, ruleID), reqOpt)
	return
}
