terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_shared_credential" "cred_password" {
  name = "Cred Password"
  type = "PASSWORD"
  username = "john-doe"
  password = "sUp3rSecr3t!"
}

resource "bamboo_shared_credential" "cred_ssh" {
  name = "Cred Ssh"
  type = "SSH"
  ssh_key = file("dummy_key")
  ssh_passphrase = "test123"
}

/*
resource "bamboo_shared_credential" "cred_aws" {
  name = "Cred Aws"
  type = "AWS"
  access_key = "test"
  secret_key = "test"
}*/
