# infra/environments/prod/main.tf
module "network" {
  source   = "../../modules/network"
  vpc_cidr = var.vpc_cidr
}

module "database" {
  source            = "../../modules/database"
  db_name           = var.db_name
  db_username       = var.db_username
  db_password       = var.db_password
  security_group_id = var.db_security_group_id
  subnet_group      = var.db_subnet_group
}