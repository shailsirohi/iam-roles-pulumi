package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func GetPolicy(roleType string) pulumi.String{
	var policy pulumi.String

	switch roleType {
	case "power-user":
		policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"PowerUser\",\"Effect\":\"Allow\",\"NotAction\":\"iam:*\",\"Resource\":\"*\"}]}"
		break
	case "eks-role":
		policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"eks-user\",\"Effect\":\"Allow\",\"Action\":\"eks:DescribeCluster\",\"Resource\":\"*\"}]}"
		break
	}

	return policy
}