# infra/modules/dns/variables.tf

variable "domain_name" {
  description = "The domain name to create DNS records for"
  type        = string
}

variable "alb_dns_name" {
  description = "The DNS name of the ALB"
  type        = string
}

variable "alb_zone_id" {
  description = "The hosted zone ID of the ALB"
  type        = string
}
