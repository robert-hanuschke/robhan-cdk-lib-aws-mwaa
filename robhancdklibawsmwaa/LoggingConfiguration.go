package robhancdklibawsmwaa


// The type of Apache Airflow logs to send to CloudWatch Logs.
type LoggingConfiguration struct {
	// Defines the processing logs sent to CloudWatch Logs and the logging level to send.
	DagProcessingLogs *ModuleLoggingConfiguration `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// Defines the scheduler logs sent to CloudWatch Logs and the logging level to send.
	SchedulerLogs *ModuleLoggingConfiguration `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// Defines the task logs sent to CloudWatch Logs and the logging level to send.
	TaskLogs *ModuleLoggingConfiguration `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// Defines the web server logs sent to CloudWatch Logs and the logging level to send.
	WebServerLogs *ModuleLoggingConfiguration `field:"optional" json:"webServerLogs" yaml:"webServerLogs"`
	// Defines the worker logs sent to CloudWatch Logs and the logging level to send.
	WorkerLogs *ModuleLoggingConfiguration `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

