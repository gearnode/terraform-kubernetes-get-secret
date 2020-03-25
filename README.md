# Terraform Module: Get Kubernetes Secret Value

> This repository is a [Terraform](https://terraform.io/) Module to use Kubernetes secrets as data.

## Table of Contents

- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Usage](#usage)
  - [Module Variables](#module-variables)
- [Contributing](#contributing)
- [Maintainers](#maintainers)

## Requirements

This module requires Terraform version `0.10.x` or newer and a valid out-of-cluster Kubernetes configuration at the default path (`$HOME/.kube/config`) or the path defined by `$KUBECONFIG`.

## Usage

Add the module to your Terraform resources:

```hcl
module "rds-postgres-password" {
  source    = "github.com/gearnode/terraform-kubernetes-get-secret?ref=v0.3.0"

  namespace = "default"
  name = "terraform"
  key = "database-password"
  context = "supercontext"
}
```

and load the module using `terraform get`.

### Module Variables

Available variables are listed below, along with their default values:

| variable    | description                      |
|-------------|----------------------------------|
| `namespace` | The kubernetes namespace         |
| `name`      | The kubernetes secret name       |
| `key`       | The kubernetes secret key to get |
| `context`   | The kubernetes context           |

### Module outputs

Available outputs are listed below, along with their description:

| output    | description                   |
|-----------|-------------------------------|
| `result`  | A string of the secret value. |

## Contributing

### Requirements

- fully installed and configured `go` environment

### Contributing code

To build binaries, use the following commands:

```
$ git clone git@github.com:gearnode/terraform-kubernetes-get-secret.git
$ cd terraform-kubernetes-get-secret
$ make
```

## Maintainers

This module is currently maintained by the individuals listed below.

- [Bryan Frimin](https://github.com/gearnode)
- [Ludovic Vielle](https://github.com/lukkor)
