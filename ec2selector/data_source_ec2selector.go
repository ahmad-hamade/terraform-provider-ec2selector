package ec2selector

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/amazon-ec2-instance-selector/v2/pkg/bytequantity"
	"github.com/aws/amazon-ec2-instance-selector/v2/pkg/selector"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInstancesRead,

		Schema: map[string]*schema.Schema{
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"vcpu": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"cpu_arch": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceInstancesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	vcpusRange := &selector.IntRangeFilter{}
	if v, ok := d.GetOk("vcpu"); ok {
		vcpusRange = &selector.IntRangeFilter{
			LowerBound: v.(int),
			UpperBound: v.(int),
		}
	} else {
		vcpusRange = nil
	}

	memoryRange := &selector.ByteQuantityRangeFilter{}
	if v, ok := d.GetOk("memory"); ok {
		memoryRange = &selector.ByteQuantityRangeFilter{
			LowerBound: bytequantity.FromGiB(v.(uint64)),
			UpperBound: bytequantity.FromGiB(v.(uint64)),
		}
	} else {
		memoryRange = nil
	}

	var cpuArch *string
	if v, ok := d.GetOk("cpu_arch"); ok {
		val := v.(string)
		cpuArch = &val
	} else {
		cpuArch = nil
	}

	// Instantiate a new instance of a selector with the AWS session
	instanceSelector := selector.New(sess)
	filters := selector.Filters{
		VCpusRange:      vcpusRange,
		MemoryRange:     memoryRange,
		CPUArchitecture: cpuArch,
	}

	// Pass the Filter struct to the Filter function of your selector instance
	instanceTypesSlice, err := instanceSelector.Filter(filters)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("instances", instanceTypesSlice); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
