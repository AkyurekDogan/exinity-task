# infra/modules/dns/outputs.tf

output "zone_id" {
  value = aws_route53_zone.this.id
  description = "The ID of the Route 53 hosted zone"
}

output "domain_name" {
  value = aws_route53_zone.this.name
  description = "The name of the domain in Route 53"
}
