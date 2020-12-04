package ec2selector

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceInstances(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEc2SelectorInstancesDataSource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ec2selector_instances.test", "id"),
				),
			},
		},
	})
}

var testAccEc2SelectorInstancesDataSource = `
data "ec2selector_instances" "test" {
  vcpu = 4
  memory = 8
}
`
