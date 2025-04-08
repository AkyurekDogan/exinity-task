module "network" {
  source = "../../modules/network"

  vpc_name     = "staging-vpc"
  subnet_count = 3
}

module "rds" {
  source = "../../modules/database"

  db_name     = "exinity_task"
  db_username = var.db_username
  db_password = var.db_password
  db_instance_class = "db.t3.small" # Use a larger instance than dev, but still small
}

module "eks" {
  source = "../../modules/compute"

  cluster_name = "staging-cluster"
  node_count   = 3
  instance_type = "t3.medium" # Medium instance for staging
}

module "iam" {
  source = "../../modules/iam"

  eks_cluster_name = "staging-cluster"
}

module "alb" {
  source = "../../modules/storage"

  alb_name = "staging-alb"
}

module "dns" {
  source = "../../modules/storage"

  domain_name = "staging.exinity-task.com"
  alb_dns_name = module.alb.dns_name
}
