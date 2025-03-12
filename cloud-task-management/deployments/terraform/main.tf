provider "aws" {
  region = "us-west-2"  # Change to your desired AWS region
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "main" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"  # Change to your desired availability zone
}

resource "aws_db_instance" "default" {
  identifier         = "task-management-db"
  engine            = "postgres"
  instance_class    = "db.t2.micro"
  allocated_storage  = 20
  username          = var.db_username
  password          = var.db_password
  db_name           = "task_management"
  skip_final_snapshot = true
  vpc_security_group_ids = [aws_security_group.db_sg.id]
  subnet_group_name = aws_db_subnet_group.default.name
}

resource "aws_db_subnet_group" "default" {
  name       = "task-management-db-subnet-group"
  subnet_ids = [aws_subnet.main.id]
}

resource "aws_security_group" "db_sg" {
  name        = "db_sg"
  description = "Allow access to the database"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"]  # Change to your desired CIDR block
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "db_endpoint" {
  value = aws_db_instance.default.endpoint
}