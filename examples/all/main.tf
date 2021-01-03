terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

# users

resource "bamboo_user" "john_doe" {
  username = "john.doe"
  full_name = "John Doe"
  email = "john@company.com"
}

resource "bamboo_user" "jane_doe" {
  username = "jane.doe"
  full_name = "Jane Doe"
  email = "jane@company.com"
}

# groups

resource "bamboo_group" "devops" {
  name = "devops"
  members = [
    bamboo_user.john_doe.username
  ]
}

resource "bamboo_group" "testers" {
  name = "testers"
  members = [
    bamboo_user.jane_doe.username
  ]
}

# global permissions

resource "bamboo_global_permission" "group_devops" {
  name = bamboo_group.devops.name
  type = "GROUP"
  permissions = [
    "READ",
    "CREATE",
    "CREATEREPOSITORY",
    "ADMINISTRATION"
  ]
}

resource "bamboo_global_permission" "group_testers" {
  name = bamboo_group.testers.name
  type = "GROUP"
  permissions = [
    "READ"
  ]
}

# shared credentials

resource "bamboo_shared_credential" "git_password" {
  name = "git-password"
  type = "PASSWORD"

  username = "john-doe"
  password = "sUp3rSecr3t!"
}

# global variables

resource "bamboo_global_variable" "docker_username" {
  key = "DockerUsername"
  value = "docker"
}

resource "bamboo_global_variable" "docker_password" {
  key = "DockerPassword"
  value = "sUp3rSecr3t!"
}

# linked repositories

resource "bamboo_linked_repository" "git_specs" {
  name = "git-specs"
  type = "GIT"

  repository_url = "https://code.company.com/scm/pub/repository.git"
  branch = "main"

  auth_type = "PASSWORD_SHARED"
  shared_credential_id = bamboo_shared_credential.git_password.id
}

# linked repository permissions

resource "bamboo_linked_repository_permission" "devops" {
  repository_id = bamboo_linked_repository.git_specs.id

  name = bamboo_group.devops.name
  type = "GROUP"
  permissions = [
    "READ",
    "ADMINISTRATION"
  ]
}

# local agents

resource "bamboo_local_agent" "agent_1" {
  name = "Agent 1"
  description = "Devops agent"
  enabled = true
}

resource "bamboo_local_agent" "agent_2" {
  name = "Agent 2"
  description = "Testers agent"
  enabled = false
}

# projects

resource "bamboo_project" "website" {
  name = "Website"
  key = "WEB"
  description = "Company website"
}
