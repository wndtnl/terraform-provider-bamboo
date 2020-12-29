package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooLocalAgent(t *testing.T) {

	rn := "bamboo_local_agent.agent"

	name := acctest.RandString(10)
	desc1 := acctest.RandString(25)
	desc2 := acctest.RandString(25)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccGlobalVariableDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccLocalAgentConfig(name, desc1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "description", desc1),
					resource.TestCheckResourceAttrSet(rn, "enabled"),
				),
			},
			{
				// Update
				Config: testAccGlobalVariableConfig(name, desc2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "description", desc2),
					resource.TestCheckResourceAttrSet(rn, "enabled"),
				),
			},
		},
	})
}

func testAccLocalAgentDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccLocalAgentConfig(name, description string) string {
	return fmt.Sprintf(`
	resource "bamboo_local_agent" "agent" {
		name = "%s"
		description = "%s"
		enabled = true
	}
	`, name, description)
}
