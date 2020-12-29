terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source = "windtunnel.io/atlassian/bamboo"
    }
  }
}

provider "bamboo" {}

resource "bamboo_local_agent" "agent_1" {
  name = "Agent 1"
  description = "Managed by Team 1"
  enabled = true
}

output "agent_id" {
  value = bamboo_local_agent.agent_1.id
}
