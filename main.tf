terraform {
  required_version = "~> 1.5.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  profile = "aws-honeypot"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway = false
  enable_vpn_gateway = false
  

  tags = {
    Terraform = "true"
    Environment = "dev"
    CreatedBy = "Umesh"
  }
}

resource "tls_private_key" "cc-kp" {
  algorithm = "RSA"
  rsa_bits  = 4096

}

resource "aws_key_pair" "generated_key" {
  
  # Name of key: Write the custom name of your key
  key_name   = "cc-key-pair"
  
  # Public Key: The public will be generated using the reference of tls_private_key.terrafrom_generated_private_key
  public_key = tls_private_key.cc-kp.public_key_openssh
 
  # Store private key :  Generate and save private key(aws_keys_pairs.pem) in current directory 
  provisioner "local-exec" {   
    command = <<-EOT
      echo '${tls_private_key.cc-kp.private_key_pem}' > aws_keys_pairs.pem
      chmod 400 aws_keys_pairs.pem
    EOT
  }
} 


module "security_group" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "my-security-group"
  description = "Security group for example usage with EC2 instance"
  vpc_id      = module.vpc.vpc_id

  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["https-443-tcp", "ssh-tcp"]
  egress_rules        = ["all-all"]
}


module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.0"

  name = "cc-instance"

  ami                    = "ami-0b0ea68c435eb488d"
  instance_type          = "t3.micro"
  monitoring             = true
  vpc_security_group_ids = [module.security_group.security_group_id]
  subnet_id              = module.vpc.public_subnets[0]
  key_name = aws_key_pair.generated_key.key_name
  // provide public ip 
  associate_public_ip_address = true
  tags = {
    Terraform = "true"
    Environment = "dev" 
    CreatedBy = "Umesh"
  }
}



output "private_key" {
  value     = tls_private_key.cc-kp.private_key_pem
  sensitive = true
}