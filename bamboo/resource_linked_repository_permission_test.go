package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccBambooLinkedRepositoryPermission(t *testing.T) {

	rn := "bamboo_linked_repository_permission.repository_permission"

	name := "admin"
	permissionType := "USER"
	permissions1 := []string{"READ"}
	permissions2 := []string{"READ", "ADMINISTRATION"}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccLinkedRepositoryPermissionDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccLinkedRepositoryPermissionConfig(name, permissionType, permissions1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "repository_id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", permissionType),
					resource.TestCheckResourceAttr(rn, "permissions.#", "1"),
					resource.TestCheckResourceAttr(rn, "permissions.0", "READ"),
				),
			},
			{
				// Update
				Config: testAccLinkedRepositoryPermissionConfig(name, permissionType, permissions2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "repository_id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", permissionType),
					resource.TestCheckResourceAttr(rn, "permissions.#", "2"),
				),
			},
		},
	})
}

func testAccLinkedRepositoryPermissionDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccLinkedRepositoryPermissionConfig(name, permissionType string, permissions []string) string {

	repoName := acctest.RandString(10)

	var p string
	if len(permissions) == 0 {
		p = "[]"
	} else {
		p = "[\"" + strings.Join(permissions, "\",\"") + "\"]"
	}

	return fmt.Sprintf(`
	resource "bamboo_linked_repository" "linked_repository" {
		name = "%s"
		type = "GIT"
		repository_url = "https://code.company.com/scm/pub/repository.git"
		branch = "master"
		auth_type = "NONE"
	}

	resource "bamboo_linked_repository_permission" "repository_permission" {
		repository_id = bamboo_linked_repository.linked_repository.id
		name = "%s"
		type = "%s"
		permissions = %s
	}
	`, repoName, name, permissionType, p)
}
