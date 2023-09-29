package softwaredeployment

import "github.com/opentelekomcloud/gophertelekomcloud"

const resourcePath = "software_deployments"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
