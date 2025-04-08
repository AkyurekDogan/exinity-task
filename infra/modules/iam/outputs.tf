# infra/modules/iam/outputs.tf

output "eks_cluster_role_arn" {
  value = aws_iam_role.eks_cluster_role.arn
  description = "The ARN of the IAM role for the EKS cluster"
}

output "eks_worker_role_arn" {
  value = aws_iam_role.eks_worker_role.arn
  description = "The ARN of the IAM role for the EKS worker nodes"
}
