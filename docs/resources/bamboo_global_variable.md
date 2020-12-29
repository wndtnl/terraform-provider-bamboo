# Bamboo Global Variable

Provides a Bamboo global variable. This can be used to create and manage Bamboo global variables.

## Example Usage

```hcl
# Create a new Bamboo global variable
resource "bamboo_global_variable" "database_password" {
  key = "DatabasePassword"
  value = "sUp3rSecr3t!"
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Required) Key for the global variable, needs to be unique and adhere to the Bamboo validation rules for variable keys.
* `value` - (Required) Value for the global variable.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the global variable.

## Import

Global variables can be imported using their key (i.e. variable name), e.g.

```
$ terraform import bamboo_global_variable.database_password DatabasePassword
```
