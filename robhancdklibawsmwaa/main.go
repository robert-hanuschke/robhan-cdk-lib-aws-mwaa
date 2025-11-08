// AWS CDK Construct Library for Amazon Managed Workflows for Apache Airflow
package robhancdklibawsmwaa

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.AirflowVersion",
		reflect.TypeOf((*AirflowVersion)(nil)).Elem(),
		map[string]interface{}{
			"V2_7_2": AirflowVersion_V2_7_2,
			"V2_8_1": AirflowVersion_V2_8_1,
			"V2_9_2": AirflowVersion_V2_9_2,
			"V2_10_1": AirflowVersion_V2_10_1,
			"V2_10_3": AirflowVersion_V2_10_3,
			"V3_0_6": AirflowVersion_V3_0_6,
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.EndpointManagement",
		reflect.TypeOf((*EndpointManagement)(nil)).Elem(),
		map[string]interface{}{
			"CUSTOMER": EndpointManagement_CUSTOMER,
			"SERVICE": EndpointManagement_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_mwaa.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptions", GoGetter: "AirflowConfigurationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersion", GoGetter: "AirflowVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "celeryExecutorQueue", GoGetter: "CeleryExecutorQueue"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3Path", GoGetter: "DagS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "databaseVpcEndpointService", GoGetter: "DatabaseVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "endpointManagement", GoGetter: "EndpointManagement"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClass", GoGetter: "EnvironmentClass"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationDagProcessingLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationSchedulerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationTaskLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationTaskLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWebserverLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWebserverLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWorkerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWorkerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "maxWebservers", GoGetter: "MaxWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "minWebservers", GoGetter: "MinWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkers", GoGetter: "MinWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersion", GoGetter: "PluginsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3Path", GoGetter: "PluginsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersion", GoGetter: "RequirementsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3Path", GoGetter: "RequirementsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "schedulers", GoGetter: "Schedulers"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucket", GoGetter: "SourceBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3ObjectVersion", GoGetter: "StartupScriptS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3Path", GoGetter: "StartupScriptS3Path"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessMode", GoGetter: "WebserverAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "webserverUrl", GoGetter: "WebserverUrl"},
			_jsii_.MemberProperty{JsiiProperty: "webserverVpcEndpointService", GoGetter: "WebserverVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStart", GoGetter: "WeeklyMaintenanceWindowStart"},
		},
		func() interface{} {
			j := jsiiProxy_Environment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EnvironmentBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentAttributes",
		reflect.TypeOf((*EnvironmentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentBase",
		reflect.TypeOf((*EnvironmentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptions", GoGetter: "AirflowConfigurationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersion", GoGetter: "AirflowVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "celeryExecutorQueue", GoGetter: "CeleryExecutorQueue"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3Path", GoGetter: "DagS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "databaseVpcEndpointService", GoGetter: "DatabaseVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "endpointManagement", GoGetter: "EndpointManagement"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClass", GoGetter: "EnvironmentClass"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationDagProcessingLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationSchedulerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationTaskLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationTaskLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWebserverLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWebserverLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWorkerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWorkerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "maxWebservers", GoGetter: "MaxWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "minWebservers", GoGetter: "MinWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkers", GoGetter: "MinWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersion", GoGetter: "PluginsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3Path", GoGetter: "PluginsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersion", GoGetter: "RequirementsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3Path", GoGetter: "RequirementsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "schedulers", GoGetter: "Schedulers"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucket", GoGetter: "SourceBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3ObjectVersion", GoGetter: "StartupScriptS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3Path", GoGetter: "StartupScriptS3Path"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessMode", GoGetter: "WebserverAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "webserverUrl", GoGetter: "WebserverUrl"},
			_jsii_.MemberProperty{JsiiProperty: "webserverVpcEndpointService", GoGetter: "WebserverVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStart", GoGetter: "WeeklyMaintenanceWindowStart"},
		},
		func() interface{} {
			j := jsiiProxy_EnvironmentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEnvironment)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentClass",
		reflect.TypeOf((*EnvironmentClass)(nil)).Elem(),
		map[string]interface{}{
			"MW1_MICRO": EnvironmentClass_MW1_MICRO,
			"MW1_SMALL": EnvironmentClass_MW1_SMALL,
			"MW1_MEDIUM": EnvironmentClass_MW1_MEDIUM,
			"MW1_LARGE": EnvironmentClass_MW1_LARGE,
			"MW1_1LARGE": EnvironmentClass_MW1_1LARGE,
			"MW1_2LARGE": EnvironmentClass_MW1_2LARGE,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_mwaa.EnvironmentProps",
		reflect.TypeOf((*EnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_mwaa.IEnvironment",
		reflect.TypeOf((*IEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptions", GoGetter: "AirflowConfigurationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersion", GoGetter: "AirflowVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "celeryExecutorQueue", GoGetter: "CeleryExecutorQueue"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3Path", GoGetter: "DagS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "databaseVpcEndpointService", GoGetter: "DatabaseVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "endpointManagement", GoGetter: "EndpointManagement"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environmentArn", GoGetter: "EnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClass", GoGetter: "EnvironmentClass"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationDagProcessingLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationSchedulerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationSchedulerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationTaskLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationTaskLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWebserverLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWebserverLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationWorkerLogsCloudWatchLogGroupArn", GoGetter: "LoggingConfigurationWorkerLogsCloudWatchLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "maxWebservers", GoGetter: "MaxWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "minWebservers", GoGetter: "MinWebservers"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkers", GoGetter: "MinWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersion", GoGetter: "PluginsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3Path", GoGetter: "PluginsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersion", GoGetter: "RequirementsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3Path", GoGetter: "RequirementsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "schedulers", GoGetter: "Schedulers"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucket", GoGetter: "SourceBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3ObjectVersion", GoGetter: "StartupScriptS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "startupScriptS3Path", GoGetter: "StartupScriptS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessMode", GoGetter: "WebserverAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "webserverUrl", GoGetter: "WebserverUrl"},
			_jsii_.MemberProperty{JsiiProperty: "webserverVpcEndpointService", GoGetter: "WebserverVpcEndpointService"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStart", GoGetter: "WeeklyMaintenanceWindowStart"},
		},
		func() interface{} {
			j := jsiiProxy_IEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"DEBUG": LogLevel_DEBUG,
			"INFO": LogLevel_INFO,
			"WARNING": LogLevel_WARNING,
			"ERROR": LogLevel_ERROR,
			"CRITICAL": LogLevel_CRITICAL,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_mwaa.LoggingConfiguration",
		reflect.TypeOf((*LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_mwaa.ModuleLoggingConfiguration",
		reflect.TypeOf((*ModuleLoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_mwaa.NetworkConfiguration",
		reflect.TypeOf((*NetworkConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.WebserverAccessMode",
		reflect.TypeOf((*WebserverAccessMode)(nil)).Elem(),
		map[string]interface{}{
			"PRIVATE_ONLY": WebserverAccessMode_PRIVATE_ONLY,
			"PUBLIC_ONLY": WebserverAccessMode_PUBLIC_ONLY,
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_mwaa.WorkerReplacementStrategy",
		reflect.TypeOf((*WorkerReplacementStrategy)(nil)).Elem(),
		map[string]interface{}{
			"FORCED": WorkerReplacementStrategy_FORCED,
			"GRACEFUL": WorkerReplacementStrategy_GRACEFUL,
		},
	)
}
