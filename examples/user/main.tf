terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_user" "john_doe" {
  username = "john.doe"
  full_name = "John Doe"
  email = "john@company.com"
  password = "john1234"
  active = true
}
