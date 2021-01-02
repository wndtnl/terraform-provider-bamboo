package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooDeploymentProject(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping testing in short mode as this test requires manual creation of a build plan")
	}

	rn := "bamboo_deployment_project.deployment_project"

	name1 := acctest.RandString(10)
	name2 := acctest.RandString(10)
	description := acctest.RandString(20)
	planKey := "WEB-COM"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccDeploymentProjectDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccDeploymentProjectConfig(name1, description, planKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name1),
					resource.TestCheckResourceAttr(rn, "description", description),
					resource.TestCheckResourceAttr(rn, "plan_key", planKey),
				),
			},
			{
				// Update
				Config: testAccDeploymentProjectConfig(name2, description, planKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name2),
					resource.TestCheckResourceAttr(rn, "description", description),
					resource.TestCheckResourceAttr(rn, "plan_key", planKey),
				),
			},
		},
	})
}

func testAccDeploymentProjectDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccDeploymentProjectConfig(name, description, planKey string) string {
	return fmt.Sprintf(`
	resource "bamboo_deployment_project" "deployment_project" {
		name = "%s"
		description = "%s"
		plan_key = "%s"
	}
	`, name, description, planKey)
}
