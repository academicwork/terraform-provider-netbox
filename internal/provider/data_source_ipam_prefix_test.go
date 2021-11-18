package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamPrefixRead(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceIpamPrefixRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.netbox_ipam_prefix.foo", "id", regexp.MustCompile(`59`)),
					resource.TestMatchResourceAttr("data.netbox_ipam_prefix.foo", "cidr", regexp.MustCompile(`10.34.0.0/21`)),
				),
			},
		},
	})
}

const testDataSourceIpamPrefixRead = `
data "netbox_ipam_prefix" "foo" {
  site = "dev-hz1"
  region = "helsinki"
}
`