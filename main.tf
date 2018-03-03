provider "aws" {
  version = "~> 1.10"
}

data "external" "secret" {
  program = ["${path.module}/bin/get-secret-value"]

  query = {
    namespace = "${var.namespace}"
    name      = "${var.name}"
    key       = "${var.key}"
  }
}
