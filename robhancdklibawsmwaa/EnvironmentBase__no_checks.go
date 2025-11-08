//go:build no_runtime_type_checking

package robhancdklibawsmwaa

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnvironmentBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EnvironmentBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EnvironmentBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEnvironmentBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironmentBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEnvironmentBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEnvironmentBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

