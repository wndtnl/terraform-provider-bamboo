# Bamboo Global Permission

Provides a Bamboo global permission for a user or group. This can be used to create and manage Bamboo global permissions.

## Example Usage

```hcl
# Create a new Bamboo group
resource "bamboo_group" "grp_devops" {
  name = "devops"
  members = [
    "admin"
  ]
}

# Create a new Bamboo global permission for a user
resource "bamboo_global_permission" "gp_admin" {
  name = "bob"
  type = "USER"
  permissions = [
    "READ",
    "CREATE",
    "CREATEREPOSITORY",
    "ADMINISTRATION"
  ]
}

# Create a new Bamboo global permission for a group
resource "bamboo_global_permission" "gp_devops" {
  name = bamboo_group.grp_devops.name
  type = "GROUP"
  permissions = [
    "READ",
    "CREATEREPOSITORY"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the user (i.e. username) or the group. Needs to be unique. Changing this forces a new resource to be created.
* `type` - (Required) Type of the global permission. Possible values are `USER` and `GROUP`. Changing this forces a new resource to be created.
* `permissions` - (Required) A list of permissions to be assigned to the user or group. Possible values are `READ` (Access), `CREATE`, `CREATEREPOSITORY` and `ADMINISTRATION`.
At least a single permission is required. Not all permission combinations are valid.

## Import

Global permissions can be imported using the format "<type>|<name>", e.g.

```
$ terraform import bamboo_global_permission.gp_devops "GROUP|grp_devops"
```
