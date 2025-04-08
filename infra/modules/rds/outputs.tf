# infra/modules/rds/outputs.tf

output "rds_instance_id" {
  value       = aws_db_instance.postgres_db.id
  description = "The ID of the RDS PostgreSQL instance"
}

output "rds_endpoint" {
  value       = aws_db_instance.postgres_db.endpoint
  description = "The endpoint of the RDS PostgreSQL instance"
}

output "rds_security_group_id" {
  value       = aws_security_group.rds_sg.id
  description = "The ID of the RDS security group"
}
