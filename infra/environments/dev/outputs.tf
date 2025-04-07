# infra/environments/dev/outputs.tf
output "db_endpoint" {
  value = module.database.db_endpoint
}

