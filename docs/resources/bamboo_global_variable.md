# Bamboo Global Variable

Provides a Bamboo global variable. This can be used to create and manage Bamboo global variables.

## Example Usage

```hcl
# Create a new Bamboo global variable
resource "bamboo_global_variable" "nexus_user" {
  key = "NexusUser"
  value = "nexus-user"
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Required) Key for the global variable, needs to be unique and adhere to the Bamboo validation rules for variable keys.
* `value` - (Required) Value for the global variable.

## Import

Global variables can be imported using their key, e.g.

```
$ terraform import bamboo_global_variable.nexus_user NexusUser
```