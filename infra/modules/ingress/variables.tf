# infra/modules/ingress/variables.tf

variable "domain_name" {
  description = "The domain name to route traffic to"
  type        = string
  default     = "www.exinity-task.com"
}

variable "service_name" {
  description = "The name of the service to route traffic to"
  type        = string
  default     = "api-service"
}

variable "service_port" {
  description = "The port the service is exposed on"
  type        = number
  default     = 80
}
