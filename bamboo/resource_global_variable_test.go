package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooGlobalVariable(t *testing.T) {

	rn := "bamboo_global_variable.variable"

	key := acctest.RandString(10)
	value1 := acctest.RandString(10)
	value2 := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		ProviderFactories:    testAccProviderFactories(),
		CheckDestroy: testAccGlobalVariableDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccGlobalVariableConfig(key, value1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "key", key),
					resource.TestCheckResourceAttr(rn, "value", value1),
				),
			},
			{
				// Update
				Config: testAccGlobalVariableConfig(key, value2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "key", key),
					resource.TestCheckResourceAttr(rn, "value", value2),
				),
			},
		},
	})
}

func testAccGlobalVariableDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccGlobalVariableConfig(key, value string) string {
	return fmt.Sprintf(`
	resource "bamboo_global_variable" "variable" {
		key = "%s"
		value = "%s"
	}
	`, key, value)
}