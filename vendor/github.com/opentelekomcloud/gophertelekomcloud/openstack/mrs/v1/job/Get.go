package job

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func Get(c *golangsdk.ServiceClient, id string) (*JobExecution, error) {
	// GET /v1.1/{project_id}/job-exes/{job_exe_id}
	raw, err := c.Get(c.ServiceURL("job-exes", id), nil, openstack.StdRequestOpts())
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*JobExecution, error) {
	if err != nil {
		return nil, err
	}

	var res JobExecution
	err = extract.IntoStructPtr(raw.Body, &res, "job_execution")
	return &res, err
}

type JobExecution struct {
	// Whether job execution objects are generated by job templates.
	Templated bool `json:"templated,omitempty"`
	// Creation time, which is a 10-bit timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// Update time, which is a 10-bit timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Job ID
	Id string `json:"id,omitempty"`
	// Project ID. For details on how to obtain the project ID
	TenantId string `json:"tenant_id,omitempty"`
	// Job application ID
	JobId string `json:"job_id,omitempty"`
	// Job name
	JobName string `json:"job_name,omitempty"`
	// Data input ID
	InputId string `json:"input_id,omitempty"`
	// Data output ID
	OutputId string `json:"output_id,omitempty"`
	// Start time of job execution, which is a 10-bit timestamp.
	StartTime int64 `json:"start_time,omitempty"`
	// End time of job execution, which is a 10-bit timestamp.
	EndTime int64 `json:"end_time,omitempty"`
	// Cluster ID
	ClusterId string `json:"cluster_id,omitempty"`
	// Workflow ID of Oozie
	EngineJobId string `json:"engine_job_id,omitempty"`
	// Returned code for an execution result
	ReturnCode string `json:"return_code,omitempty"`
	// Whether a job is public
	// The current version does not support this function.
	IsPublic bool `json:"is_public,omitempty"`
	// Whether a job is protected
	// The current version does not support this function.
	IsProtected bool `json:"is_protected,omitempty"`
	// Group ID of a job
	GroupId string `json:"group_id,omitempty"`
	// Path of the .jar file for program execution
	JarPath string `json:"jar_path,omitempty"`
	// Address for inputting data
	Input string `json:"input,omitempty"`
	// Address for outputting data
	Output string `json:"output,omitempty"`
	// Address for storing job logs
	JobLog string `json:"job_log,omitempty"`
	// Job type code
	// 1: MapReduce
	// 2: Spark
	// 3: Hive Script
	// 4: HiveQL (not supported currently)
	// 5: DistCp
	// 6: Spark Script
	// 7: Spark SQL (not supported in this API currently)
	JobType int `json:"job_type,omitempty"`
	// Data import and export
	FileAction string `json:"file_action,omitempty"`
	// Key parameter for program execution. The parameter is specified by the function of the user's internal program.
	// MRS is only responsible for loading the parameter. This parameter can be empty.
	Arguments string `json:"arguments,omitempty"`
	// Job status code
	// -1: Terminated
	// 1: Starting
	// 2: Running
	// 3: Completed
	// 4: Abnormal
	// 5: Error
	JobState int `json:"job_state,omitempty"`
	// Final job status
	// 0: unfinished
	// 1: terminated due to an execution error
	// 2: executed successfully
	// 3: canceled
	JobFinalStatus int `json:"job_final_status,omitempty"`
	// Address of the Hive script
	HiveScriptPath string `json:"hive_script_path,omitempty"`
	// User ID for creating jobs
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	CreateBy string `json:"create_by,omitempty"`
	// Number of completed steps
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	FinishedStep int `json:"finished_step,omitempty"`
	// Main ID of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	JobMainId string `json:"job_main_id,omitempty"`
	// Step ID of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	JobStepId string `json:"job_step_id,omitempty"`
	// Delay time, which is a 10-bit timestamp.
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	PostponeAt int64 `json:"postpone_at,omitempty"`
	// Step name of a job
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	StepName string `json:"step_name,omitempty"`
	// Number of steps
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	StepNum int `json:"step_num,omitempty"`
	// Number of tasks
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	TaskNum int `json:"task_num,omitempty"`
	// User ID for updating jobs
	UpdateBy string `json:"update_by,omitempty"`
	// Token
	// The current version does not support this function.
	Credentials string `json:"credentials,omitempty"`
	// User ID for creating jobs
	// This parameter is not used in the current version, but is retained for compatibility with earlier versions.
	UserId string `json:"user_id,omitempty"`
	// Key-value pair set for saving job running configurations
	JobConfigs map[string]interface{} `json:"job_configs,omitempty"`
	// Authentication information
	// The current version does not support this function.
	Extra map[string]interface{} `json:"extra,omitempty"`
	// Data source URL
	DataSourceUrls map[string]interface{} `json:"data_source_urls,omitempty"`
	// Key-value pair set, containing job running information returned by Oozie
	Info map[string]interface{} `json:"info,omitempty"`
	// Encrypt Type
	EncryptType int `json:"encrypt_type,omitempty"`
}
