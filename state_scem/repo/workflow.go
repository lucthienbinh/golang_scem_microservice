package repo

// ----------------- TABLE IN DATABASE -----------------

// WorkflowModel structure
type WorkflowModel struct {
	ID                     uint   `gorm:"primary_key;<-:false" json:"id"`
	WorkflowVersion        int    `json:"workflow_version"`
	WorkflowKey            string `json:"workflow_key"`
	Step                   int    `json:"step"`
	Type                   int    `json:"type"`
	Name                   uint   `json:"name"`
	NextStep1              int    `json:"next_step_1"`
	NextStep2              int    `json:"next_step_2"`
	ServiceRetry           int    `json:"service_retry"`
	MessageCorrelationName string `json:"message_correlation_name"`
}

// WorkflowInstance structure
type WorkflowInstance struct {
	ID              uint   `gorm:"primary_key;<-:false" json:"id"`
	WorkflowVersion int    `json:"workflow_version"`
	WorkflowKey     string `json:"workflow_key"`
	CurrentStep     int    `json:"current_step"`
	CurrentType     int    `json:"current_type"`
	CurrentName     uint   `json:"current_name"`
	Running         bool   `json:"running"`
	Finished        bool   `json:"finished"`
	Failed          bool   `json:"failed"`
	Canceled        bool   `json:"canceled"`
}

// WorkflowRuningPath structure
type WorkflowRuningPath struct {
	ID                 uint  `gorm:"primary_key;<-:false" json:"id"`
	WorkflowInstanceID uint  `json:"workflow_instance_id"`
	Step               int   `json:"step"`
	Type               int   `json:"type"`
	Name               uint  `json:"name"`
	FinishedTime       int64 `json:"finished_time"`
}

// WorkflowVariable structure
type WorkflowVariable struct {
	ID                 uint   `gorm:"primary_key;<-:false" json:"id"`
	WorkflowInstanceID uint   `json:"workflow_instance_id"`
	VariableName       string `json:"variable_name"`
	VariableValue      string `json:"variable_value"`
}

// WorkflowJobQueue structure
type WorkflowJobQueue struct {
	ID                 uint   `gorm:"primary_key;<-:false" json:"id"`
	WorkflowInstanceID uint   `json:"workflow_instance_id"`
	Name               string `json:"name"`
	RetryRemain        int    `json:"retry_remain"`
	Finished           bool   `json:"finished"`
	Failed             bool   `json:"failed"`
	Canceled           bool   `json:"canceled"`
}

// WorkflowMessageQueue structure
type WorkflowMessageQueue struct {
	ID                      uint   `json:"id"`
	WorkflowInstanceID      uint   `json:"workflow_instance_id"`
	Name                    uint   `json:"name"`
	MessageCorrelationName  string `json:"message_correlation_name"`
	MessageCorrelationValue int    `json:"message_correlation_value"`
	Finished                bool   `json:"finished"`
	Failed                  bool   `json:"failed"`
}