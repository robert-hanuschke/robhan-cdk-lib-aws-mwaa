//go:build no_runtime_type_checking

package robhancdklibawsmwaa

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Environment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEnvironment_FromEnvironmentAttributesParameters(scope constructs.Construct, id *string, attrs *EnvironmentAttributes) error {
	return nil
}

func validateEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironment_IsEnvironmentParameters(x interface{}) error {
	return nil
}

func validateEnvironment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEnvironment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEnvironmentParameters(scope constructs.Construct, id *string, props *EnvironmentProps) error {
	return nil
}

