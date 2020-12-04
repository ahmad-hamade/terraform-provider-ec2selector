terraform {
  required_providers {
    ec2selector = {
      source = "hamade.me/edu/ec2selector"
      versions = "=0.1.0"
    }
  }
}

provider "ec2selector" {}

data "ec2selector_instances" "all" {
  vcpu = 2
  memory = 1
}

output "ec2" {
  value = data.ec2selector_instances.all.instances
}
