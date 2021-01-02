terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_linked_repository" "git" {
  name = "Company Repo"
  type = "GIT"

  repository_url = "https://code.company.com/scm/pub/repository.git"
  branch = "main"

  auth_type = "NONE"
}

resource "bamboo_linked_repository_permission" "admin" {

  repository_id = bamboo_linked_repository.git.id

  name = "admin"
  type = "USER"
  permissions = [
    "READ",
    "ADMINISTRATION"
  ]
}

resource "bamboo_linked_repository_permission" "bamboo-admin" {

  repository_id = bamboo_linked_repository.git.id

  name = "bamboo-admin"
  type = "GROUP"
  permissions = [
    "READ",
    "ADMINISTRATION"
  ]
}
