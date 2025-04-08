output "rds_endpoint" {
  description = "The endpoint of the RDS database"
  value       = module.rds.db_endpoint
}

output "eks_cluster_url" {
  description = "The URL of the EKS cluster"
  value       = module.eks.cluster_url
}

output "alb_dns_name" {
  description = "The DNS name of the ALB"
  value       = module.alb.dns_name
}
