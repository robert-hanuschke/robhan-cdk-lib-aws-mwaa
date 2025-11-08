package robhancdklibawsmwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/robert-hanuschke/robhan-cdk-lib-aws-mwaa/robhancdklibawsmwaa/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-mwaa/robhancdklibawsmwaa/internal"
)

type EnvironmentBase interface {
	awscdk.Resource
	IEnvironment
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	//
	// For example, core.default_timezone: utc.
	AirflowConfigurationOptions() *map[string]*string
	// The version of Apache Airflow to use for the environment.
	//
	// If no value is specified, defaults to the latest version.
	//
	// If you specify a newer version number for an existing environment, the version update requires some service interruption before taking effect.
	AirflowVersion() AirflowVersion
	// The queue ARN for the environment's Celery Executor.
	//
	// Amazon MWAA uses a Celery Executor to distribute tasks across multiple workers.
	// When you create an environment in a shared VPC, you must provide access to the Celery Executor queue from your VPC.
	CeleryExecutorQueue() *string
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	//
	// For example, dags.
	DagS3Path() *string
	// The VPC endpoint for the environment's Amazon RDS database.
	DatabaseVpcEndpointService() *string
	// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
	//
	// If set to SERVICE, Amazon MWAA will create and manage the required VPC endpoints in your VPC.
	// If set to CUSTOMER, you must create, and manage, the VPC endpoints in your VPC.
	EndpointManagement() EndpointManagement
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN for the Amazon MWAA environment.
	EnvironmentArn() *string
	// The environment class type.
	EnvironmentClass() EnvironmentClass
	// The execution role in IAM that allows MWAA to access AWS resources in your environment.
	ExecutionRole() awsiam.IRole
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	//
	// You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).
	KmsKey() awskms.IKey
	// The Apache Airflow logs being sent to CloudWatch Logs.
	LoggingConfiguration() *LoggingConfiguration
	// The ARN for the CloudWatch Logs group where the Apache Airflow DAG processing logs are published.
	LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Scheduler logs are published.
	LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow task logs are published.
	LoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Web server logs are published.
	LoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Worker logs are published.
	LoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string
	// The maximum number of web servers that you want to run in your environment.
	//
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for MaxWebservers when you interact with your Apache
	// Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. For example, in scenarios where your workload requires network
	// calls to the Apache Airflow REST API with a high transaction-per-second (TPS) rate, Amazon MWAA will increase the number of web servers up to
	// the number set in MaxWebserers. As TPS rates decrease Amazon MWAA disposes of the additional web servers, and scales down to the number set in
	// MinxWebserers.
	//
	// Valid values: For environments larger than mw1.micro, accepts values from 2 to 5. Defaults to 2 for all environment sizes except mw1.micro,
	// which defaults to 1.
	MaxWebservers() *float64
	// The maximum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the MaxWorkers field. For example, 20. When there are no more
	// tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or
	// the number you specify in MinWorkers.
	MaxWorkers() *float64
	// The minimum number of web servers that you want to run in your environment.
	//
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for MaxWebservers when you interact with your Apache
	// Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. As the transaction-per-second rate, and the network load,
	// decrease, Amazon MWAA disposes of the additional web servers, and scales down to the number set in MinxWebserers.
	//
	// Valid values: For environments larger than mw1.micro, accepts values from 2 to 5. Defaults to 2 for all environment sizes except mw1.micro,
	// which defaults to 1.
	MinWebservers() *float64
	// The minimum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you
	// specify in the MaxWorkers field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving
	// the worker count you specify in the MinWorkers field. For example, 2.
	MinWorkers() *float64
	// The name of your Amazon MWAA environment.
	Name() *string
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	NetworkConfiguration() *NetworkConfiguration
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The version of the plugins.zip file on your Amazon S3 bucket.
	PluginsS3ObjectVersion() *string
	// The relative path to the plugins.zip file on your Amazon S3 bucket. For example, plugins.zip.
	PluginsS3Path() *string
	// The version of the requirements.txt file on your Amazon S3 bucket.
	RequirementsS3ObjectVersion() *string
	// The relative path to the requirements.txt file on your Amazon S3 bucket. For example, requirements.txt.
	RequirementsS3Path() *string
	// The number of schedulers that you want to run in your environment.
	//
	// Valid values:
	// v2 - For environments larger than mw1.micro, accepts values from 2 to 5.
	//      Defaults to 2 for all environment sizes except mw1.micro, which defaults to 1.
	// v1 - Accepts 1.
	Schedulers() *float64
	// The Amazon S3 bucket where your DAG code and supporting files are stored.
	SourceBucket() awss3.IBucket
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The version of the startup shell script in your Amazon S3 bucket.
	//
	// You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	// Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than 1,024 bytes long.
	//
	// The following is an example:
	// 3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo.
	StartupScriptS3ObjectVersion() *string
	// The relative path to the startup shell script in your Amazon S3 bucket.
	//
	// For example, s3://mwaa-environment/startup.sh.
	// Amazon MWAA runs the script as your environment starts, and before running the Apache Airflow process.
	// You can use this script to install dependencies, modify Apache Airflow configuration options, and set environment variables.
	StartupScriptS3Path() *string
	// The Apache Airflow Web server access mode.
	WebserverAccessMode() WebserverAccessMode
	// The URL of your Apache Airflow UI.
	WebserverUrl() *string
	// The VPC endpoint for the environment's web server.
	WebserverVpcEndpointService() *string
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: DAY:HH:MM.
	//
	// For example: TUE:03:30. You can specify a start time in 30 minute increments only.
	//
	// Supported input includes the following:
	// MON|TUE|WED|THU|FRI|SAT|SUN:([01]\\d|2[0-3]):(00|30).
	WeeklyMaintenanceWindowStart() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EnvironmentBase
type jsiiProxy_EnvironmentBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IEnvironment
}

func (j *jsiiProxy_EnvironmentBase) AirflowConfigurationOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) AirflowVersion() AirflowVersion {
	var returns AirflowVersion
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) CeleryExecutorQueue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"celeryExecutorQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) DatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) EndpointManagement() EndpointManagement {
	var returns EndpointManagement
	_jsii_.Get(
		j,
		"endpointManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) EnvironmentClass() EnvironmentClass {
	var returns EnvironmentClass
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfiguration() *LoggingConfiguration {
	var returns *LoggingConfiguration
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationDagProcessingLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationSchedulerLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationTaskLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationWebserverLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) LoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationWorkerLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) MaxWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) MinWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) NetworkConfiguration() *NetworkConfiguration {
	var returns *NetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) SourceBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"sourceBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) WebserverAccessMode() WebserverAccessMode {
	var returns WebserverAccessMode
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) WebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentBase) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}


func NewEnvironmentBase_Override(e EnvironmentBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentBase",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EnvironmentBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnvironmentBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func EnvironmentBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EnvironmentBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_EnvironmentBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := e.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := e.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

