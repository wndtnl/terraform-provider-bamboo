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

  auth_type = "PASSWORD"
  username = "john.doe"
  password = "sUp3rS3crEt!"

  submodules = true
  lfs = true
}

output "repository_id" {
  value = bamboo_linked_repository.git.id
}
