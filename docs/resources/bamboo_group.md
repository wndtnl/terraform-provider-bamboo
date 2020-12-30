# Bamboo Group

Provides a Bamboo (user) group. This can be used to create and manage Bamboo groups.

## Example Usage

```hcl
# Create a new Bamboo group
resource "bamboo_group" "grp_devops" {
  name = "devops"
  members = [
    "admin"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the group, needs to be unique. Changing this forces a new resource to be created.
* `members` - (Optional) A list of group members referenced by username. An error is returned when the provided username does not match an existing user.

## Import

Groups can be imported using their name, e.g.

```
$ terraform import bamboo_group.grp_devops "devops"
```
