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

resource "bamboo_global_permission" "gp_admin" {
  name = "bob"
  type = "USER"
  permissions = [
    "READ",
    "CREATE",
    "CREATEREPOSITORY",
    "ADMINISTRATION"
  ]
}

resource "bamboo_global_permission" "gp_devops" {
  name = bamboo_group.grp_devops.name
  type = "GROUP"
  permissions = [
    "READ",
    "CREATEREPOSITORY"
  ]
}
