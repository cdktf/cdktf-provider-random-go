// Prebuilt random Provider for Terraform CDK (cdktf)
package random

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The maximum inclusive value of the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/integer#max Integer#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// The minimum inclusive value of the range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/integer#min Integer#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource.
	//
	// See [the main provider documentation](../index.html) for more information.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/integer#keepers Integer#keepers}
	Keepers *map[string]*string `field:"optional" json:"keepers" yaml:"keepers"`
	// A custom seed to always produce the same value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/integer#seed Integer#seed}
	Seed *string `field:"optional" json:"seed" yaml:"seed"`
}

