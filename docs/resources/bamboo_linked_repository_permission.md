# Bamboo Linked Repository Permission

Provides a Bamboo linked repository permission for a user or group. This can be used to create and manage Bamboo repository permissions.

## Example Usage

```hcl
# Create a new Bamboo linked repository
resource "bamboo_linked_repository" "git" {
  name = "Company Repo"
  type = "GIT"

  repository_url = "https://code.company.com/scm/pub/repository.git"
  branch = "main"

  auth_type = "NONE"
}

# Create a new Bamboo group
resource "bamboo_group" "grp_devops" {
  name = "devops"
  members = [
    "admin"
  ]
}

# Create a new Bamboo linked repository permission for a user
resource "bamboo_linked_repository_permission" "rp_admin" {

  repository_id = bamboo_linked_repository.git.id

  name = "admin"
  type = "USER"
  permissions = [
    "READ"
  ]
}

# Create a new Bamboo linked repository permission for a group
resource "bamboo_linked_repository_permission" "rp_devops" {

  repository_id = bamboo_linked_repository.git.id
  
  name = bamboo_group.grp_devops.name
  type = "GROUP"
  permissions = [
    "READ",
    "ADMINISTRATION"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) ID of the linked repository. Changing this forces a new resource to be created.
* `name` - (Required) Name of the user (i.e. username) or the group. Needs to be unique. Changing this forces a new resource to be created.
* `type` - (Required) Type of the permission. Possible values are `USER` and `GROUP`. Changing this forces a new resource to be created.
* `permissions` - (Required) A list of permissions to be assigned to the user or group. Possible values are `READ` (Access) and `ADMINISTRATION`.
  At least a single permission is required. Not all permission combinations are valid.

## Import

Linked repository permissions can be imported using the format "<repository-id>|<type>|<name>", e.g.

```
$ terraform import bamboo_linked_repository_permission.rp_devops "6389770|GROUP|grp_devops"
```
