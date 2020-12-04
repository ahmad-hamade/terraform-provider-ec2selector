# Terraform Provider ec2selector

This is a terraform provider that lets you get a list of EC2 instances type based on vcpu and memory along with other filters.
Its based on [amazon-ec2-instance-selector](https://github.com/aws/amazon-ec2-instance-selector) liblary.

Run the following command to build the provider

```shell
go build -o terraform-provider-ec2selector
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd examples
terraform init && terraform apply
```
