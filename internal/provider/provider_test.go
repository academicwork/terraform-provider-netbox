package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"netbox": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = New("dev")()
	testAccProviders = map[string]*schema.Provider{
		"netbox": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("NETBOX_HOST"); err == "" {
		t.Fatal("NETBOX_HOST must be set for acceptance tests")
	}
	if err := os.Getenv("NETBOX_TOKEN"); err == "" {
		t.Fatal("NETBOX_TOKEN must be set for acceptance tests")
	}
}
