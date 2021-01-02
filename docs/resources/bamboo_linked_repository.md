# Bamboo Linked Repository

Provides a Bamboo linked repository. This can be used to create and manage Bamboo linked repositories.

> Only linked repositories of type 'GIT' are supported at this point.

## Example Usage

```hcl
# Create a new Bamboo linked repository
resource "bamboo_linked_repository" "git" {

  name = "Company Repo"
  type = "GIT"

  repository_url = "https://code.company.com/scm/pub/repository.git"
  branch = "main"

  auth_type = "PASSWORD"
  username = "john.doe"
  password = "sUp3rS3crEt!"

  submodules = true
  lfs = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the linked repository, needs to be unique.
* `type` - (Required) Type of the linked repository. Possible values are `GIT`.
* `repository_url` - (Required) Url of the git repository. Must have a valid form depending on the selected `auth_type`.
* `branch` - (Optional) Branch of the git repository. Defaults to `master` if not specified.
* `auth_type` - (Required) The repository authentication type. Possible values are `NONE`, `PASSWORD`, `SSH`, `PASSWORD_SHARED` and `SSH_SHARED`.
* `username` - (Optional) This must be specified when `auth_type` is set to `PASSWORD`.
* `password` - (Optional) This must be specified when `auth_type` is set to `PASSWORD`.
* `ssh_key` - (Optional) This must be specified when `auth_type` is set to `SSH`.
* `ssh_passphrase` - (Optional) This can be specified when `auth_type` is set to `SSH`.
* `shared_credential_id` - (Optional) This must be specified when `auth_type` is set to `PASSWORD_SHARED` or `SSH_SHARED`.
* `shallow_clones` - (Optional) Use shallow clones. Defaults to `false` if not specified.
* `remote_agent_cache` - (Optional) Enable repository caching on remote agents. Defaults to `true` if not specified.
* `submodules` - (Optional) Use submodules. Defaults to `false` if not specified.
* `verbose_logs` - (Optional) Verbose logs. Defaults to `false` if not specified.
* `fetch_whole_repository` - (Optional) Fetch whole repository. Defaults to `false` if not specified.
* `lfs` - (Optional) Enable LFS support. Defaults to `false` if not specified.
* `command_timeout` - (Optional) Command timeout in minutes. Defaults to `180` if not specified.
* `quit_period` - (Optional) Enable quiet period. Defaults to `false` if not specified.
* `quiet_period_wait_time` - (Optional) Quiet period in seconds. Defaults to `10` if not specified.
* `quiet_period_max_retries` - (Optional) Maximum retries. Defaults to `5` if not specified.
* `filter_pattern` - (Optional) Include / exclude file mode. Possible values are `NONE`, `INCLUDE_ONLY` and `EXCLUDE_ALL`.
* `filter_pattern_regex` - (Optional) File pattern regular expression.
* `change_set_regex` - (Optional) Exclude changesets regular expression.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the linked repository.

## Import

Linked repositories can be imported using their name, e.g.

```
$ terraform import bamboo_linked_repository.git "Company Repo"
```
