# Bamboo Project

Provides a Bamboo project. This can be used to create and manage Bamboo projects.

> Deletion of a project might take up to a few minutes because projects are first marked for deletion by Bamboo,
> after which they are asynchronously removed from the database in a background task. The provider waits till the
> project is fully deleted as to ensure the project name and key are available again.

> Take note that deleting a project will also delete any child plans.

## Example Usage

```hcl
# Create a new Bamboo project
resource "bamboo_project" "website" {
  name = "Website"
  key = "WEB"
  description = "Company website"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the project, needs to be unique.
* `key` - (Required) Key for the project. 
  Must have a minimum length of two (2) characters, must start with a capital letter and can contain capital letters (A-Z) and numbers (0-9).
  Changing this forces a new resource to be created.
* `description` - (Optional) Description for the project.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the project.

## Import

Projects can be imported using their key, e.g.

```
$ terraform import bamboo_project.website WEB
```
