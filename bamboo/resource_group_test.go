package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccBambooGroup(t *testing.T) {

	rn := "bamboo_group.group"

	name := acctest.RandString(10)
	member := "admin"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccGroupDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccGroupConfig(name, []string{member}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "members.#", "1"),
					resource.TestCheckResourceAttr(rn, "members.0", member),
				),
			},
			{
				// Update
				Config: testAccGroupConfig(name, []string{}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "members.#", "0"),
				),
			},
		},
	})
}

func testAccGroupDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccGroupConfig(name string, members []string) string {

	var s string
	if len(members) == 0 {
		s = "[]"
	} else {
		s = "[\"" + strings.Join(members, "\",\"") + "\"]"
	}

	return fmt.Sprintf(`
	resource "bamboo_group" "group" {
		name = "%s"
		members = %s
	}
	`, name, s)
}
