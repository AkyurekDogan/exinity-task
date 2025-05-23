module "network" {
  source = "../../modules/network"

  vpc_name     = "dev-vpc"
  subnet_count = 3
}

module "rds" {
  source = "../../modules/database"

  db_name     = "exinity_task"
  db_username = var.db_username
  db_password = var.db_password
  db_instance_class = "db.t3.micro" # Use a smaller instance for dev
}

module "eks" {
  source = "../../modules/compute"

  cluster_name = "dev-cluster"
  node_count   = 2
  instance_type = "t3.medium" # A smaller instance for dev
}

module "iam" {
  source = "../../modules/iam"

  eks_cluster_name = "dev-cluster"
}

module "alb" {
  source = "../../modules/storage"

  alb_name = "dev-alb"
}

module "dns" {
  source = "../../modules/storage"

  domain_name = "dev.exinity-task.com"
  alb_dns_name = module.alb.dns_name
}
