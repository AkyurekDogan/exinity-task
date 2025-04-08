# infra/modules/iam/variables.tf

variable "cluster_role_name" {
  description = "The name of the IAM role for the EKS cluster"
  type        = string
}

variable "worker_role_name" {
  description = "The name of the IAM role for the EKS worker nodes"
  type        = string
}
