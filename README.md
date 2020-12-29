# Terraform Provider for Bamboo

Terraform provider for the "Terraform Provider for Bamboo" Atlassian plugin.

To update the go-bamboo client sdk:

```
$ go get -u github.com/wndtnl/go-bamboo@<commit>
```

Running examples:

```
$ make
$ cd examples/...

$ terraform init

$ terraform apply -state="terraform.tfstate"
$ terraform destroy -auto-approve -state="terraform.tfstate"
```

Creating a release:

```
export GITHUB_TOKEN="..."

git tag -a v0.0.1
git push origin v0.0.1
```

which will trigger the github release action.
