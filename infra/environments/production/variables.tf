variable "db_username" {
  description = "The username for the PostgreSQL database"
  type        = string
  default     = "prod_postgres_user"
}

variable "db_password" {
  description = "The password for the PostgreSQL database"
  type        = string
  sensitive   = true
  default     = "prod_password_123"
}

variable "instance_type" {
  description = "Type of EC2 instances to use in EKS"
  type        = string
  default     = "m5.large"
}

variable "node_count" {
  description = "The number of nodes in the EKS cluster"
  type        = number
  default     = 6
}

variable "vpc_name" {
  description = "Name of the VPC"
  type        = string
  default     = "prod-vpc"
}

variable "subnet_count" {
  description = "Number of subnets to create"
  type        = number
  default     = 3
}
