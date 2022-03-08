package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx,"")
		roleType := conf.Require("roleType")
		name := conf.Require("roleName")

		//Get appropriate policy
		policy := GetPolicy(roleType)

		// Create an AWS resource (S3 Bucket)
		iamRole, err := iam.NewRole(ctx,name,&iam.RoleArgs{
			Name: pulumi.String(name),
			InlinePolicies: iam.RoleInlinePolicyArray{
				iam.RoleInlinePolicyArgs{
					Name: pulumi.String("access-policy"),
					Policy: policy,
				},
			},
			AssumeRolePolicy: pulumi.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"),
		})
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", iamRole.Arn)
		return nil
	})
}