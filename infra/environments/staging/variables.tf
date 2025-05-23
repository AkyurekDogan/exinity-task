variable "db_username" {
  description = "The username for the PostgreSQL database"
  type        = string
  default     = "staging_postgres_user"
}

variable "db_password" {
  description = "The password for the PostgreSQL database"
  type        = string
  sensitive   = true
  default     = "staging_password_123"
}

variable "instance_type" {
  description = "Type of EC2 instances to use in EKS"
  type        = string
  default     = "t3.medium"
}

variable "node_count" {
  description = "The number of nodes in the EKS cluster"
  type        = number
  default     = 3
}

variable "vpc_name" {
  description = "Name of the VPC"
  type        = string
  default     = "staging-vpc"
}

variable "subnet_count" {
  description = "Number of subnets to create"
  type        = number
  default     = 3
}
