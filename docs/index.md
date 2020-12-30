# Bamboo Provider

This [Bamboo](https://www.atlassian.com/software/bamboo) provider can be used to interact with Bamboo resources
as exposed by the [Terraform Provider for Bamboo](https://windtunnel.io/products/tpb) add-on. It is important to understand that:

- This provider has not been created nor is it maintained by [Atlassian](https://www.atlassian.com) but by [WindTunnel Technologies](https://windtunnel.io),
an independent Marketplace vendor.
- This provider allows managing Bamboo resources which are not exposed through the [official Bamboo API](https://docs.atlassian.com/atlassian-bamboo/REST/latest/)
which makes the installation of the accompanying [Marketplace add-on](https://marketplace.atlassian.com) in the Bamboo server a mandatory prerequisite.

-> Please refer to the [add-on documentation](https://windtunnel.io/products/tpb) for additional installation and getting started instructions.
Any questions or remarks can be directed to [hello@windtunnel.io](mailto:hello@windtunnel.io).

## Example Usage

```hcl
# Required for Terraform 0.13 and up (https://www.terraform.io/upgrade-guides/0-13.html)
terraform {
  required_version = ">=0.13"
  required_providers {
    bamboo = {
      source  = "registry.terraform.io/wndtnl/bamboo"
      version = "0.0.1"
    }
  }
}

# Configure the Bamboo provider
provider "bamboo" {
  address = "${var.bamboo_address}"
  username = "${var.bamboo_username}"
  password = "${var.bamboo_password}"
}

# Create a global variable
resource "bamboo_global_variable" "nexus_user" {
  key = "NexusUser"
  value = "bamboo"
}
```

## Authentication

The provider currently supports a single means of authentication, Basic Auth.

### Basic Auth

Add the `username` and `password` field to the provider block. Alternatively, these values can be provided
through the environment variables `BAMBOO_USER` and `BAMBOO_PASS`.

Usage:

```hcl
# Configure the Bamboo provider
provider "bamboo" {
  address = "http://localhost:6990/bamboo"
  username = "admin"
  password = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Optional) The base address of the Bamboo server. This value can also be provided using the `BAMBOO_ADDR` environment variable.
When no value can be found, the provider falls back to a default `http://localhost:6990/bamboo` value.
* `username` - (Optional) The username of an administrative user. This value can also be provided using the `BAMBOO_USER` environment variable.
When no value can be found, the provider falls back to a default `admin` value.
* `password` - (Optional) The password of an administrative user. This value can also be provided using the `BAMBOO_PASS` environment variable.
When no value can be found, the provider falls back to a default `admin` value.
