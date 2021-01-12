package repo

import "context"

// Repository interface
type Repository interface {
	// WorkflowModel
	CreateWorkflowModel(ctx context.Context, workflowModel WorkflowModel) (uint, bool, error)
	GetWorkflowModelLastestVersionByProcessID(ctx context.Context, processID string) (WorkflowModel, bool, error)

	// WorkflowInstance
	CreateWorkflowInstance(ctx context.Context, workflowInstance WorkflowInstance) (uint, bool, error)
	GetWorkflowInstanceList(ctx context.Context) ([]WorkflowInstance, bool, error)
	GetWorkflowInstance(ctx context.Context, id uint) (WorkflowInstance, bool, error)
	UpdateWorkflowInstance(ctx context.Context, id uint, workflowInstance WorkflowInstance) (bool, error)

	// WorkflowRuningPath
	CreateWorkflowRuningPath(ctx context.Context, workflowRuningPath WorkflowRuningPath) (uint, bool, error)
	GetWorkflowRuningPathList(ctx context.Context) ([]WorkflowRuningPath, bool, error)
	GetWorkflowRuningPathListByWFInstanceID(ctx context.Context, workflowInstanceID uint) ([]WorkflowRuningPath, bool, error)
	UpdateWorkflowRuningPath(ctx context.Context, id uint, workflowRuningPath WorkflowRuningPath) (bool, error)

	// WorkflowVariable
	CreateWorkflowVariable(ctx context.Context, workflowVariable WorkflowVariable) (uint, bool, error)
	GetWorkflowVariableList(ctx context.Context) ([]WorkflowVariable, bool, error)
	GetWorkflowVariableListByWFInstanceID(ctx context.Context, workflowInstanceID uint) ([]WorkflowVariable, bool, error)
	UpdateWorkflowVariable(ctx context.Context, id uint, workflowVariable WorkflowVariable) (bool, error)

	// WorkflowJobQueue
	CreateWorkflowJobQueue(ctx context.Context, workflowJobQueue WorkflowJobQueue) (uint, bool, error)
	GetWorkflowJobQueueList(ctx context.Context) ([]WorkflowJobQueue, bool, error)
	GetWorkflowJobQueueListByMappingName(ctx context.Context, mappingName string) ([]WorkflowJobQueue, bool, error)
	UpdateWorkflowJobQueue(ctx context.Context, id uint, workflowJobQueue WorkflowJobQueue) (bool, error)

	// WorkflowMessageQueue
	CreateWorkflowMessageQueue(ctx context.Context, workflowMessageQueue WorkflowMessageQueue) (uint, bool, error)
	GetWorkflowMessageQueueList(ctx context.Context) ([]WorkflowMessageQueue, bool, error)
	GetWorkflowMessageQueueListByMappingName(ctx context.Context, mappingName string) ([]WorkflowMessageQueue, bool, error)
	UpdateWorkflowMessageQueue(ctx context.Context, id uint, workflowMessageQueue WorkflowMessageQueue) (bool, error)
}

// InitRepository interface
type InitRepository interface {
	DeleteDatabase(ctx context.Context) error
	MigrationDatabase(ctx context.Context) error
}
