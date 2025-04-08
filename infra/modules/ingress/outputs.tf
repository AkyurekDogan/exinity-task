# infra/modules/ingress/outputs.tf

output "api_ingress_name" {
  value = kubernetes_ingress.api_ingress.metadata[0].name
  description = "The name of the API Ingress resource"
}

output "api_ingress_host" {
  value = var.domain_name
  description = "The domain name for the API"
}
