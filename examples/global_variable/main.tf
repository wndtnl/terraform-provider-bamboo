terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_global_variable" "nexus_user" {
  key = "NexusUser2"
  value = "nexus-user2"
}

resource "bamboo_global_variable" "database_password" {
  key = "DatabasePassword"
  value = "sUp3rSecr3t!"
}

output "db_password_id" {
  value = bamboo_global_variable.database_password.id
}
