# Bamboo Local Agent

Provides a Bamboo local agent. This can be used to create and manage Bamboo local agents.

## Example Usage

```hcl
# Create a new Bamboo local agent
resource "bamboo_local_agent" "agent_1" {
  name = "Agent 1"
  description = "Managed by Team 1"
  enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the local agent, needs to be unique.
* `description` - (Optional) Description for the local agent.
* `enabled` - (Optional) Set the agent status. Defaults to `true` if not specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the local agent.

## Import

Local agents can be imported using their name, e.g.

```
$ terraform import bamboo_local_agent.agent_1 "Agent 1"
```
