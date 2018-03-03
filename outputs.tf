output "result" {
  value = "${chomp(base64decode(data.external.secret.result["value"]))}"
}
