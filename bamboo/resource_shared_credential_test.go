package bamboo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccBambooSharedCredential_Password(t *testing.T) {

	rn := "bamboo_shared_credential.credential_password"

	name := acctest.RandString(10)
	username := acctest.RandString(15)
	password1 := acctest.RandString(15)
	password2 := acctest.RandString(15)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccSharedCredentialDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccSharedCredentialPasswordConfig(name, username, password1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", "PASSWORD"),
					resource.TestCheckResourceAttr(rn, "username", username),
					resource.TestCheckResourceAttr(rn, "password", password1),
				),
			},
			{
				// Update
				Config: testAccSharedCredentialPasswordConfig(name, username, password2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", "PASSWORD"),
					resource.TestCheckResourceAttr(rn, "username", username),
					resource.TestCheckResourceAttr(rn, "password", password2),
				),
			},
		},
	})
}

func TestAccBambooSharedCredential_Ssh(t *testing.T) {

	rn := "bamboo_shared_credential.credential_ssh"

	name := acctest.RandString(10)
	sshKey := acctest.RandString(200)
	sshPassphrase1 := acctest.RandString(15)
	sshPassphrase2 := acctest.RandString(15)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		CheckDestroy:      testAccSharedCredentialDestroy,
		Steps: []resource.TestStep{
			{
				// Create
				Config: testAccSharedCredentialSshConfig(name, sshKey, sshPassphrase1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", "SSH"),
					resource.TestCheckResourceAttr(rn, "ssh_key", sshKey),
					resource.TestCheckResourceAttr(rn, "ssh_passphrase", sshPassphrase1),
				),
			},
			{
				// Update
				Config: testAccSharedCredentialSshConfig(name, sshKey, sshPassphrase2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "id"),
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "type", "SSH"),
					resource.TestCheckResourceAttr(rn, "ssh_key", sshKey),
					resource.TestCheckResourceAttr(rn, "ssh_passphrase", sshPassphrase2),
				),
			},
		},
	})
}

func testAccSharedCredentialDestroy(s *terraform.State) error {
	return nil // NOP
}

func testAccSharedCredentialPasswordConfig(name, username, password string) string {
	return fmt.Sprintf(`
	resource "bamboo_shared_credential" "credential_password" {
		name = "%s"
		type = "PASSWORD"
		username = "%s"
		password = "%s"
	}
	`, name, username, password)
}

func testAccSharedCredentialSshConfig(name, key, passphrase string) string {
	return fmt.Sprintf(`
	resource "bamboo_shared_credential" "credential_ssh" {
		name = "%s"
		type = "SSH"
		ssh_key = "%s"
		ssh_passphrase = "%s"
	}
	`, name, key, passphrase)
}

func testAccSharedCredentialAwsConfig(name, accessKey, secretKey string) string {
	return fmt.Sprintf(`
	resource "bamboo_shared_credential" "credential_aws" {
		name = "%s"
		type = "AWS"
		access_key = "%s"
		secret_key = "%s"
	}
	`, name, accessKey, secretKey)
}
