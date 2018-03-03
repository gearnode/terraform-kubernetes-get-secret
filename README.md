# Terraform Module: Get Kubernetes Secret Value
> This repository is a [Terraform](https://terraform.io/) Module for help to manage secret value management in terraform, via Kubernetes Secret Store Object.

## Table of Contents

- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Usage](#usage)
  - [Module Variables](#module-variables)
- [Maintainers](#maintainers)

## Requirements

This module requires Terraform version `0.10.x` or newer.

## Dependencies

This module depends on a correctly configured and connected [kubectl binary](https://kubernetes.io/docs/tasks/tools/install-kubectl/) in the host system. This module depends too, on a properly installed [jq binary](https://stedolan.github.io/jq/) in the host system.

## Usage

Add the module to your Terraform resources like so:

```hcl
module "rds-postgres-password" {
  source    = "github.com/gearnode/terraform-kubernetes-get-secret?ref=0.1.0"
  namespace = "default"
  key       = "terraform"
  name      = "database-password"
}
```

Then, load the module using `terraform get`.

### Module Variables

Available variables are listed below, along with their default values:

| variable                                | description |
|-----------------------------------------|-|
| `namespace`                             | The kubernetes namespace |
| `name`                                  | The kubernetes secret name |
| `key`                                   | The kubernetes secret key to get |

### Module outputs

Available outputs are listed below, along with their description:

| output    | description |
|-----------|-|
| `result`  | A string of the secret value. |

## Maintainers

This module is currently maintained by the individuals listed below.

- [Bryan Frimin](https://github.com/gearnode)
