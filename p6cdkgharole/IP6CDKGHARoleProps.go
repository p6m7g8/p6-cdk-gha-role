package p6cdkgharole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type IP6CDKGHARoleProps interface {
	Policies() *[]awsiam.IManagedPolicy
	Principle() awscdk.Arn
	Repo() *string
}

// The jsii proxy for IP6CDKGHARoleProps
type jsiiProxy_IP6CDKGHARoleProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IP6CDKGHARoleProps) Policies() *[]awsiam.IManagedPolicy {
	var returns *[]awsiam.IManagedPolicy
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IP6CDKGHARoleProps) Principle() awscdk.Arn {
	var returns awscdk.Arn
	_jsii_.Get(
		j,
		"principle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IP6CDKGHARoleProps) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

