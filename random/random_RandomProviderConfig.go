// Prebuilt random Provider for Terraform CDK (cdktf)
package random


type RandomProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random#alias RandomProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

