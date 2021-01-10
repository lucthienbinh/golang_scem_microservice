package repo

// Repository interface
type Repository interface {
	// WorkflowModel
	CreateWorkflowModel(workflowModels WorkflowModel) (uint, bool, string)

	// WorkflowInstance
	CreateWorkflowInstance(workflowKey string, workflowVersion int) (uint, bool, string)
	GetWorkflowInstanceList() ([]WorkflowInstance, bool, string)
	GetWorkflowInstance(id uint) (WorkflowInstance, bool, string)
	UpdateWorkflowInstance(id uint, workflowInstance WorkflowInstance) (bool, string)

	// WorkflowRuningPath
	CreateWorkflowRuningPath(workflowRuningPath WorkflowRuningPath) (uint, bool, string)
	GetWorkflowRuningPathList() ([]WorkflowRuningPath, bool, string)
	GetWorkflowRuningPathListByWFInstanceID(workflowInstanceID uint) ([]WorkflowRuningPath, bool, string)
	UpdateWorkflowRuningPath(id uint, workflowRuningPath WorkflowRuningPath) (bool, string)

	// WorkflowVariable
	CreateWorkflowVariable(workflowVariable WorkflowVariable) (uint, bool, string)
	GetWorkflowVariableList() ([]WorkflowVariable, bool, string)
	GetWorkflowVariableListByWFInstanceID(workflowInstanceID uint) ([]WorkflowVariable, bool, string)
	UpdateWorkflowVariable(id uint, workflowVariable WorkflowVariable) (bool, string)

	// WorkflowJobQueue
	CreateWorkflowJobQueue(workflowJobQueue WorkflowJobQueue) (uint, bool, string)
	GetWorkflowJobQueueList() ([]WorkflowJobQueue, bool, string)
	GetWorkflowJobQueueListByName(name string) ([]WorkflowJobQueue, bool, string)
	UpdateWorkflowJobQueue(id uint, workflowJobQueue WorkflowJobQueue) (bool, string)

	// WorkflowMessageQueue
	CreateWorkflowMessageQueue(workflowMessageQueue WorkflowMessageQueue) (uint, bool, string)
	GetWorkflowMessageQueueList() ([]WorkflowMessageQueue, bool, string)
	GetWorkflowMessageQueueListbyName(name string) ([]WorkflowMessageQueue, bool, string)
	UpdateWorkflowMessageQueue(id uint, workflowMessageQueue WorkflowMessageQueue) (bool, string)
}
