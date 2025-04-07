# infra/environments/staging/variables.tf
variable "vpc_cidr" {
  default = "10.0.0.0/16"
}

variable "db_name" {
  default = "exinity_task"
}

variable "db_username" {
  default = "postgres"
}

variable "db_password" {
  default = "mypassword123!"
}

variable "db_security_group_id" {
  description = "Security group ID for the DB"
  type        = string
}

variable "db_subnet_group" {
  description = "Subnet group for the DB"
  type        = string
}