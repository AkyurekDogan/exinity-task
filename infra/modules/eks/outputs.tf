# infra/modules/eks/outputs.tf

output "eks_cluster_name" {
  value = aws_eks_cluster.this.name
  description = "The name of the EKS cluster"
}

output "eks_cluster_endpoint" {
  value = aws_eks_cluster.this.endpoint
  description = "The endpoint of the EKS cluster"
}

output "eks_cluster_arn" {
  value = aws_eks_cluster.this.arn
  description = "The ARN of the EKS cluster"
}
