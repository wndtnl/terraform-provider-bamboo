package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooProject(t *testing.T) {

	rn := "bamboo_project.project"

	name1 := acctest.RandString(10)
	name2 := acctest.RandString(10)
	key := "PRJ"
	description := acctest.RandString(20)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccProjectConfig(name1, key, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name1),
					resource.TestCheckResourceAttr(rn, "key", key),
					resource.TestCheckResourceAttr(rn, "description", description),
				),
			},
			{
				// Update
				Config: testAccProjectConfig(name2, key, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name2),
					resource.TestCheckResourceAttr(rn, "key", key),
					resource.TestCheckResourceAttr(rn, "description", description),
				),
			},
		},
	})
}

func testAccProjectDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccProjectConfig(name, key, description string) string {
	return fmt.Sprintf(`
	resource "bamboo_project" "project" {
		name = "%s"		
		key = "%s"
		description = "%s"
	}
	`, name, key, description)
}
