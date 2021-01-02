package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooLinkedRepository(t *testing.T) {

	rn := "bamboo_linked_repository.repository"

	name1 := acctest.RandString(10)
	name2 := acctest.RandString(10)

	url := "https://code.company.com/scm/pub/repository.git"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccLinkedRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccLinkedRepositoryConfig(name1, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name1),
					resource.TestCheckResourceAttr(rn, "type", "GIT"),
					resource.TestCheckResourceAttr(rn, "repository_url", url),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "auth_type", "PASSWORD"),
					resource.TestCheckResourceAttr(rn, "username", "username"),
					resource.TestCheckResourceAttr(rn, "password", "password"),
					resource.TestCheckResourceAttr(rn, "shallow_clones", "true"),
					resource.TestCheckResourceAttr(rn, "remote_agent_cache", "true"),
					resource.TestCheckResourceAttr(rn, "submodules", "true"),
					resource.TestCheckResourceAttr(rn, "verbose_logs", "true"),
					resource.TestCheckResourceAttr(rn, "fetch_whole_repository", "true"),
					resource.TestCheckResourceAttr(rn, "lfs", "true"),
					resource.TestCheckResourceAttr(rn, "command_timeout", "180"),
					resource.TestCheckResourceAttr(rn, "quit_period", "true"),
					resource.TestCheckResourceAttr(rn, "quiet_period_wait_time", "20"),
					resource.TestCheckResourceAttr(rn, "quiet_period_max_retries", "10"),
					resource.TestCheckResourceAttr(rn, "filter_pattern", "INCLUDE_ONLY"),
					resource.TestCheckResourceAttr(rn, "filter_pattern_regex", "git"),
					resource.TestCheckResourceAttr(rn, "change_set_regex", "txt"),
				),
			},
			{
				// Update
				Config: testAccLinkedRepositoryConfig(name2, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name2),
				),
			},
		},
	})
}

func testAccLinkedRepositoryDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccLinkedRepositoryConfig(name, repositoryUrl string) string {
	return fmt.Sprintf(`
	resource "bamboo_linked_repository" "repository" {
		name = "%s"
		type = "GIT"
		repository_url = "%s"
		branch = "master"
		auth_type = "PASSWORD"
		username = "username"
		password = "password"
		shallow_clones = true
		remote_agent_cache = true
		submodules = true
		verbose_logs = true
		fetch_whole_repository = true
		lfs = true
		command_timeout = 180
		quit_period = true
		quiet_period_wait_time = 20
		quiet_period_max_retries = 10
		filter_pattern = "INCLUDE_ONLY"
		filter_pattern_regex = "git"
		change_set_regex = "txt"
	}
	`, name, repositoryUrl)
}
