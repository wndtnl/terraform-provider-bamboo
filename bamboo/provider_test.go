package bamboo

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

var testAccProvider *schema.Provider
var testAccProviderFactories func() map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = func() map[string]func() (*schema.Provider, error) {
		return map[string]func() (*schema.Provider, error){
			"bamboo": func() (*schema.Provider, error) {
				return testAccProvider, nil
			},
		}
	}
}

func testAccPreCheck(t *testing.T) {
	// NOP
}