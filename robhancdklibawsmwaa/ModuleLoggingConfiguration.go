package robhancdklibawsmwaa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Defines the type of logs to send for the Apache Airflow log type (e.g. DagProcessingLogs).
type ModuleLoggingConfiguration struct {
	// The CloudWatch Logs log group for each type ofApache Airflow log type that you have enabled.
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// Indicates whether to enable the Apache Airflow log type (e.g. DagProcessingLogs) in CloudWatch Logs.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Defines the Apache Airflow logs to send for the log type (e.g. DagProcessingLogs) to CloudWatch Logs.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
}

