package archivepolicies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List makes a request against the Gnocchi API to list archive policies.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return ArchivePolicyPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieves a specific Gnocchi archive policy based on its name.
func Get(c *gophercloud.ServiceClient, archivePolicyName string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, archivePolicyName), &r.Body, nil)
	return
}
