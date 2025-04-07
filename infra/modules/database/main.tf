# infra/modules/database/main.tf
resource "aws_db_instance" "postgres" {
  identifier             = "exinity-postgres"
  engine                 = "postgres"
  instance_class         = "db.t3.micro"
  name                   = var.db_name
  username               = var.db_username
  password               = var.db_password
  allocated_storage      = 20
  storage_type           = "gp2"
  publicly_accessible    = false
  skip_final_snapshot    = true
  vpc_security_group_ids = [var.security_group_id]
  db_subnet_group_name   = var.subnet_group
}
