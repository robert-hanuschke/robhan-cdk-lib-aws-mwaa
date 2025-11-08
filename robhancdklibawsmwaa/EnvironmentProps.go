package robhancdklibawsmwaa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for creating an Amazon Managed Workflows for Apache Airflow Environment.
type EnvironmentProps struct {
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	//
	// For example, core.default_timezone: utc.
	AirflowConfigurationOptions *map[string]*string `field:"required" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// The name of your Amazon MWAA environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of Apache Airflow to use for the environment.
	//
	// If no value is specified, defaults to the latest version.
	//
	// If you specify a newer version number for an existing environment, the version update requires some service interruption before taking effect.
	AirflowVersion AirflowVersion `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	//
	// For example, dags.
	DagS3Path *string `field:"optional" json:"dagS3Path" yaml:"dagS3Path"`
	// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
	//
	// If set to SERVICE, Amazon MWAA will create and manage the required VPC endpoints in your VPC.
	// If set to CUSTOMER, you must create, and manage, the VPC endpoints in your VPC.
	EndpointManagement EndpointManagement `field:"optional" json:"endpointManagement" yaml:"endpointManagement"`
	// The environment class type.
	EnvironmentClass EnvironmentClass `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// The execution role in IAM that allows MWAA to access AWS resources in your environment.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	//
	// You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Apache Airflow logs being sent to CloudWatch Logs.
	LoggingConfiguration *LoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
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
	MaxWebservers *float64 `field:"optional" json:"maxWebservers" yaml:"maxWebservers"`
	// The maximum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the MaxWorkers field. For example, 20. When there are no more
	// tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or
	// the number you specify in MinWorkers.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The minimum number of web servers that you want to run in your environment.
	//
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for MaxWebservers when you interact with your Apache
	// Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. As the transaction-per-second rate, and the network load,
	// decrease, Amazon MWAA disposes of the additional web servers, and scales down to the number set in MinxWebserers.
	//
	// Valid values: For environments larger than mw1.micro, accepts values from 2 to 5. Defaults to 2 for all environment sizes except mw1.micro,
	// which defaults to 1.
	MinWebservers *float64 `field:"optional" json:"minWebservers" yaml:"minWebservers"`
	// The minimum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you
	// specify in the MaxWorkers field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving
	// the worker count you specify in the MinWorkers field. For example, 2.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	NetworkConfiguration *NetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The version of the plugins.zip file on your Amazon S3 bucket.
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// The relative path to the plugins.zip file on your Amazon S3 bucket. For example, plugins.zip.
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// The version of the requirements.txt file on your Amazon S3 bucket.
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// The relative path to the requirements.txt file on your Amazon S3 bucket. For example, requirements.txt.
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// The number of schedulers that you want to run in your environment.
	//
	// Valid values:
	// v2 - For environments larger than mw1.micro, accepts values from 2 to 5.
	//      Defaults to 2 for all environment sizes except mw1.micro, which defaults to 1.
	// v1 - Accepts 1.
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// The Amazon S3 bucket where your DAG code and supporting files are stored.
	SourceBucket awss3.IBucket `field:"optional" json:"sourceBucket" yaml:"sourceBucket"`
	// The version of the startup shell script in your Amazon S3 bucket.
	//
	// You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	// Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than 1,024 bytes long.
	//
	// The following is an example:
	// 3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo.
	StartupScriptS3ObjectVersion *string `field:"optional" json:"startupScriptS3ObjectVersion" yaml:"startupScriptS3ObjectVersion"`
	// The relative path to the startup shell script in your Amazon S3 bucket.
	//
	// For example, s3://mwaa-environment/startup.sh.
	// Amazon MWAA runs the script as your environment starts, and before running the Apache Airflow process.
	// You can use this script to install dependencies, modify Apache Airflow configuration options, and set environment variables.
	StartupScriptS3Path *string `field:"optional" json:"startupScriptS3Path" yaml:"startupScriptS3Path"`
	// The Apache Airflow Web server access mode.
	WebserverAccessMode WebserverAccessMode `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: DAY:HH:MM.
	//
	// For example: TUE:03:30. You can specify a start time in 30 minute increments only.
	//
	// Supported input includes the following:
	// MON|TUE|WED|THU|FRI|SAT|SUN:([01]\\d|2[0-3]):(00|30).
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

