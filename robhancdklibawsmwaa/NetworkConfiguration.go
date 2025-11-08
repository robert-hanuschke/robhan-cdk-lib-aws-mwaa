package robhancdklibawsmwaa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
type NetworkConfiguration struct {
	// A list of one or more security groups.
	//
	// Accepts up to 5 security groups. A security group must be attached to the same VPC as the subnets.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of subnets.
	//
	// Required to create an environment. Must be private subnets in two different availability zones.
	// A subnet must be attached to the same VPC as the security group.
	Subnets *[]awsec2.ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
}

