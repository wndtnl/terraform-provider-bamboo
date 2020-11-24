terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

//data "bamboo_global_variables" "variables" {}

resource "bamboo_global_variable" "repo-user" {
  key = "RepoUser2"
  value = "nexus-user"
}

resource "bamboo_global_variable" "database" {
  key = "Database4"
  value = "Connection4"
}

//output "variables" {
//  value = data.bamboo_global_variables.variables.global_variables
//}

# output "repo-user" {
#   value = bamboo_global_variable.repo-user
# }