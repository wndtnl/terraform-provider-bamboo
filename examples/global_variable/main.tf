terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

data "bamboo_global_variables" "variables" {}

resource "bamboo_global_variable" "nexus_user" {
  key = "NexusUser"
  value = "nexus-user"
}

resource "bamboo_global_variable" "database" {
  key = "Database"
  value = "Server=myServerAddress;Database=myDataBase;Uid=myUsername;Pwd=myPassword;"
}

output "variables" {
  value = data.bamboo_global_variables.variables.global_variables
}
