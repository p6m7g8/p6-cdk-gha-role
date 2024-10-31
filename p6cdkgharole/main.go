// AWS CDK: An IAM Role for Github Actions to Assume
package p6cdkgharole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"p6-cdk-gha-role.IP6CDKGHARoleProps",
		reflect.TypeOf((*IP6CDKGHARoleProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberProperty{JsiiProperty: "principle", GoGetter: "Principle"},
			_jsii_.MemberProperty{JsiiProperty: "repo", GoGetter: "Repo"},
		},
		func() interface{} {
			return &jsiiProxy_IP6CDKGHARoleProps{}
		},
	)
	_jsii_.RegisterClass(
		"p6-cdk-gha-role.P6CDKGHARole",
		reflect.TypeOf((*P6CDKGHARole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_P6CDKGHARole{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
}
