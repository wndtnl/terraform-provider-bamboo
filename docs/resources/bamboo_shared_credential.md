# Bamboo Shared Credential

Provides a Bamboo shared credential. This can be used to create and manage Bamboo shared credentials.

## Example Usage

```hcl
# Create a new Bamboo PASSWORD shared credential
resource "bamboo_shared_credential" "cred_password" {
  name = "Cred Password"
  type = "PASSWORD"
  username = "john-doe"
  password = "sUp3rSecr3t!"
}

# Create a new Bamboo SSH shared credential
resource "bamboo_shared_credential" "cred_ssh" {
  name = "Cred Ssh"
  type = "SSH"
  ssh_key = file("id_rsa")
  ssh_passphrase = "test123"
}

# Create a new Bamboo AWS shared credential
resource "bamboo_shared_credential" "cred_aws" {
  name = "Cred Aws"
  type = "AWS"
  access_key = "tbNlogn2wW2soOxeTAjF"
  secret_key = "8JL0Y3CY7tNqLdSN0A5lzpfxf/4WDaY1ljZQyAex"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the shared credential, needs to be unique.
* `type` - (Required) Type of the shared credential. Possible values are `PASSWORD`, `SSH` and `AWS`.
* `username` - (Optional) This must be specified when `type` is set to `PASSWORD`.
* `password` - (Optional) This must be specified when `type` is set to `PASSWORD`.
* `ssh_key` - (Optional) This must be specified when `type` is set to `SSH`.
* `ssh_passphrase` - (Optional) This can be specified when `type` is set to `SSH`.
* `access_key` - (Optional) This must be specified when `type` is set to `AWS`.
* `secret_key` - (Optional) This must be specified when `type` is set to `AWS`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the shared credential.

## Import

Shared credentials can be imported using their name, e.g.

```
$ terraform import bamboo_shared_credential.cred_password "Cred Password"
```
