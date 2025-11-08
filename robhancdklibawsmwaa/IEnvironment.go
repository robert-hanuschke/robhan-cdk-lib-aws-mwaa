package robhancdklibawsmwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-mwaa/robhancdklibawsmwaa/internal"
)

type IEnvironment interface {
	awscdk.IResource
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
}

// The jsii proxy for IEnvironment
type jsiiProxy_IEnvironment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEnvironment) AirflowConfigurationOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) AirflowVersion() AirflowVersion {
	var returns AirflowVersion
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) CeleryExecutorQueue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"celeryExecutorQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) DatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) EndpointManagement() EndpointManagement {
	var returns EndpointManagement
	_jsii_.Get(
		j,
		"endpointManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) EnvironmentClass() EnvironmentClass {
	var returns EnvironmentClass
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfiguration() *LoggingConfiguration {
	var returns *LoggingConfiguration
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationDagProcessingLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationSchedulerLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationTaskLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationWebserverLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) LoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingConfigurationWorkerLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) MaxWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) MinWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) NetworkConfiguration() *NetworkConfiguration {
	var returns *NetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) SourceBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"sourceBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) WebserverAccessMode() WebserverAccessMode {
	var returns WebserverAccessMode
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) WebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}

