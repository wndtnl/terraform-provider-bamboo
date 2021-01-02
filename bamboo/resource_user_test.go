package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooUser(t *testing.T) {

	rn := "bamboo_user.user"

	username := acctest.RandString(10)
	fullName1 := acctest.RandString(15)
	fullName2 := acctest.RandString(15)
	email := fmt.Sprintf("%s@company.com", username)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccUserDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccUserConfig(username, fullName1, email, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "username", username),
					resource.TestCheckResourceAttr(rn, "full_name", fullName1),
					resource.TestCheckResourceAttr(rn, "email", email),
					resource.TestCheckResourceAttr(rn, "active", "true"),
				),
			},
			{
				// Update
				Config: testAccUserConfig(username, fullName2, email, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "username", username),
					resource.TestCheckResourceAttr(rn, "full_name", fullName2),
					resource.TestCheckResourceAttr(rn, "email", email),
					resource.TestCheckResourceAttr(rn, "active", "false"),
				),
			},
		},
	})
}

func testAccUserDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccUserConfig(username, fullName, email string, active bool) string {
	return fmt.Sprintf(`
	resource "bamboo_user" "user" {
		username = "%s"
		full_name = "%s"
		email = "%s"
		active = %t
	}
	`, username, fullName, email, active)
}
