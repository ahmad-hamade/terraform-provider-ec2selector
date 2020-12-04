terraform {
  required_providers {
    ec2selector = {
      versions = "~> 0.0.1"
      source = "hamade.me/edu/ec2selector"
    }
  }
}

provider "ec2selector" {}

data "ec2selector_instances" "all" {
  vcpu = 4
  memory = 8
}

output "ec2" {
  value = data.ec2selector_instances.all.instances
}
