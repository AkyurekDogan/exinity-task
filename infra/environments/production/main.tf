module "network" {
  source = "../../modules/network"

  vpc_name     = "prod-vpc"
  subnet_count = 3
}

module "rds" {
  source = "../../modules/database"

  db_name     = "exinity_task"
  db_username = var.db_username
  db_password = var.db_password
  db_instance_class = "db.m5.large" # Larger RDS instance for production
}

module "eks" {
  source = "../../modules/compute"

  cluster_name = "prod-cluster"
  node_count   = 6
  instance_type = "m5.large" # Larger instance type for production
}

module "iam" {
  source = "../../modules/iam"

  eks_cluster_name = "prod-cluster"
}

module "alb" {
  source = "../../modules/storage"

  alb_name = "prod-alb"
}

module "dns" {
  source = "../../modules/storage"

  domain_name = "www.exinity-task.com"
  alb_dns_name = module.alb.dns_name
}
