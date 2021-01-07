package service

import "context"

// Repository interface
type Repository interface {
	Create(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

// WorkflowModel structure
type WorkflowModel struct {
	ID                     uint   `json:"id"`
	WorkflowKey            int    `json:"workflow_key"`
	Step                   int    `json:"step"`
	Type                   int    `json:"type"`
	Name                   uint   `json:"name"`
	NextStep1              int    `json:"next_step_1"`
	NextStep2              int    `json:"next_step_2"`
	ServiceRetry           int    `json:"service_retry"`
	MessageCorrelationName string `json:"message_correlation_name"`
}

// WorkflowRuningPath structure
type WorkflowRuningPath struct {
	ID                 uint  `json:"id"`
	WorkflowInstanceID uint  `json:"workflow_instance_id"`
	Step               int   `json:"step"`
	Type               int   `json:"type"`
	Name               uint  `json:"name"`
	FinishedTime       int64 `json:"finished_time"`
}

// WorkflowInstance structure
type WorkflowInstance struct {
	ID          uint `json:"id"`
	WorkflowKey int  `json:"workflow_key"`
	CurrentStep int  `json:"current_step"`
	CurrentType int  `json:"current_type"`
	CurrentName uint `json:"current_name"`
	Running     bool `json:"running"`
	Finished    bool `json:"finished"`
	Failed      bool `json:"failed"`
	Canceled    bool `json:"canceled"`
}

// WorkflowVariable structure
type WorkflowVariable struct {
	ID                 uint   `json:"id"`
	WorkflowInstanceID uint   `json:"workflow_instance_id"`
	VariableName       string `json:"variable_name"`
	VariableValue      string `json:"variable_value"`
}

// WorkflowServicePool structure
type WorkflowServicePool struct {
	ID                 uint `json:"id"`
	WorkflowInstanceID uint `json:"workflow_instance_id"`
	ServiceName        uint `json:"service_name"`
	RetryRemain        int  `json:"retry_remain"`
	Finished           bool `json:"finished"`
	Failed             bool `json:"failed"`
	Canceled           bool `json:"canceled"`
}

// WorkflowMessagePool structure
type WorkflowMessagePool struct {
	ID                      uint   `json:"id"`
	WorkflowInstanceID      uint   `json:"workflow_instance_id"`
	MessageName             uint   `json:"message_name"`
	MessageCorrelationName  string `json:"message_correlation_name"`
	MessageCorrelationValue int    `json:"message_correlation_value"`
}
