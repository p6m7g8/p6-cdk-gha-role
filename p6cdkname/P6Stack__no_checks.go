//go:build no_runtime_type_checking

package p6cdkname

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_P6Stack) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_P6Stack) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_P6Stack) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateP6Stack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateP6Stack_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateP6Stack_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewP6StackParameters(scope constructs.Construct, id *string, _props IP6Props) error {
	return nil
}

