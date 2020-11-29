# Bamboo Global Variables

Provides a Bamboo global variables datasource. This can be used to fetch all global variables from Bamboo.

## Example Usage

```hcl
# 
data "bamboo_global_variables" "variables" {}

output "variables" {
  value = data.bamboo_global_variables.variables.global_variables
}
```

## Attribute Reference

For each variable in the provided list, the following attributes are exported:

* `id` - The internal Bamboo id of the global variable.
* `key` - The key of the global variable.
* `value` - The value of the global variable.