package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccBambooGlobalPermission(t *testing.T) {

	rn := "bamboo_global_permission.global_permission"

	name := "admin"
	permissionType := "USER"
	permissions1 := []string{"READ"}
	permissions2 := []string{"READ", "CREATE"}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccGlobalPermissionDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccGlobalPermissionConfig(name, permissionType, permissions1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", permissionType),
					resource.TestCheckResourceAttr(rn, "permissions.#", "1"),
					resource.TestCheckResourceAttr(rn, "permissions.0", "READ"),
				),
			},
			{
				// Update
				Config: testAccGlobalPermissionConfig(name, permissionType, permissions2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", permissionType),
					resource.TestCheckResourceAttr(rn, "permissions.#", "2"),
				),
			},
		},
	})
}

func testAccGlobalPermissionDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccGlobalPermissionConfig(name, permissionType string, permissions []string) string {

	var p string
	if len(permissions) == 0 {
		p = "[]"
	} else {
		p = "[\"" + strings.Join(permissions, "\",\"") + "\"]"
	}

	return fmt.Sprintf(`
	resource "bamboo_global_permission" "global_permission" {
		name = "%s"
		type = "%s"
		permissions = %s
	}
	`, name, permissionType, p)
}
