terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_project" "website" {
  name = "Website"
  key = "WEB"
  description = "Company website"
}

resource "bamboo_deployment_project" "website" {
  name = "Website"
  description = "Company website"
  plan_key = "${bamboo_project.website.key}-WEB"
}
