# infra/modules/rds/variables.tf

variable "vpc_id" {
  description = "The VPC ID where the RDS instance will be deployed"
  type        = string
}

variable "db_cidr_blocks" {
  description = "The CIDR blocks that are allowed to access the database"
  type        = list(string)
  default     = ["0.0.0.0/0"]  # You should adjust this for better security
}

variable "db_allocated_storage" {
  description = "The allocated storage size (in GB) for the RDS instance"
  type        = number
  default     = 20
}

variable "db_storage_type" {
  description = "The storage type for the RDS instance"
  type        = string
  default     = "gp2"
}

variable "db_instance_class" {
  description = "The instance class for the RDS instance"
  type        = string
  default     = "db.t3.micro"
}

variable "db_name" {
  description = "The name of the PostgreSQL database"
  type        = string
  default     = "exinity_task"
}

variable "db_username" {
  description = "The username for the PostgreSQL database"
  type        = string
}

variable "db_password" {
  description = "The password for the PostgreSQL database"
  type        = string
  sensitive   = true
}

variable "subnet_ids" {
  description = "The subnet IDs where the RDS instance will be deployed"
  type        = list(string)
}
