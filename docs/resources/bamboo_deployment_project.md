# Bamboo Deployment Project

Provides a Bamboo deployment project. This can be used to create and manage Bamboo deployment projects.

> Take note that deleting a deployment project will also delete any child environments.

## Example Usage

```hcl
# Create a new Bamboo project
resource "bamboo_project" "website" {
  name = "Website"
  key = "WEB"
  description = "Company website"
}

# Create a new Bamboo deployment project
resource "bamboo_deployment_project" "website" {
  name = "Website"
  description = "Company website"
  plan_key = "${bamboo_project.website.key}-WEB"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the deployment project, needs to be unique.
* `description` - (Optional) Description for the project.
* `plan_key` - (Required) Key of the linked build plan. Can use an optional branch index suffix (e.g. PRJ-PLN3).

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the deployment project.

## Import

Deployment projects can be imported using their id, e.g.

```
$ terraform import bamboo_deployment_project.website 5832705
```
