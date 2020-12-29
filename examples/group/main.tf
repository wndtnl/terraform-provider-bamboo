terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_group" "grp_devops" {
  name = "devops"
  members = [
    "admin"
  ]
}
