# infra/modules/rds/main.tf

# Create RDS Security Group
resource "aws_security_group" "rds_sg" {
  name        = "rds-security-group"
  description = "Security group for RDS PostgreSQL"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = var.db_cidr_blocks
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "RDS Security Group"
  }
}

# Create RDS PostgreSQL instance
resource "aws_db_instance" "postgres_db" {
  allocated_storage    = var.db_allocated_storage
  storage_type         = var.db_storage_type
  engine               = "postgres"
  engine_version       = "13.3"
  instance_class       = var.db_instance_class
  db_name              = var.db_name
  username             = var.db_username
  password             = var.db_password
  db_subnet_group_name = aws_db_subnet_group.main_db_subnet_group.name
  vpc_security_group_ids = [
    aws_security_group.rds_sg.id
  ]
  multi_az             = false
  publicly_accessible  = true
  skip_final_snapshot  = true
  tags = {
    Name = "PostgreSQL RDS Instance"
  }
}

# Create DB Subnet Group
resource "aws_db_subnet_group" "main_db_subnet_group" {
  name       = "main-db-subnet-group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "Main DB Subnet Group"
  }
}

# Create DB Parameter Group
resource "aws_db_parameter_group" "postgresql_parameter_group" {
  name   = "postgresql-parameter-group"
  family = "postgres13"

  parameters = {
    "log_statement" = "all"
  }

  tags = {
    Name = "PostgreSQL Parameter Group"
  }
}
