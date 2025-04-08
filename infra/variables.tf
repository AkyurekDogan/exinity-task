variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-east-1"
}

variable "project_name" {
  description = "The name of the project"
  type        = string
  default     = "exinity-task"
}

variable "environment" {
  description = "The environment name (e.g. dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "domain_name" {
  description = "The domain name to be used for public access"
  type        = string
  default     = "www.exinity-task.com"
}