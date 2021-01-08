package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
)

var errRepo = errors.New("Unable to handle Repo Request")

// GRPCRepository interface
type GRPCRepository interface {
	DeployWorkflowModel(workflowModels []WorkflowModel) (int, error)
	CreateWorkflowInstance(workflowKey string, workflowVersion int, workflowVariable []WorkflowVariable) (uint, error)
	// Job queue
	PollingJobWorker(jobName string) (uint, []WorkflowVariable)
	CompleteJob(jobQueueID uint) error
	FailJob(jobQueueID uint) error
	// Message queue
	PublishMessage(messageName, messageCorrelationName, messageCorrelationValue string) error
}

type grpcRepo struct {
	db     *gorm.DB
	logger log.Logger
}

// NewGRPCRepo function
func NewGRPCRepo(db *gorm.DB, logger log.Logger) GRPCRepository {
	return &grpcRepo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *grpcRepo) DeployWorkflowModel(workflowModels []WorkflowModel) error {
	workflowKey := workflowModels[0].WorkflowKey
	workflowModel := &WorkflowModel{}
	version := int(1)
	if err := repo.db.Table("workflow_models as w").Select("w.workflow_version").Last(workflowModel, "w.workflow_key == ?", workflowKey).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	version = workflowModel.WorkflowVersion + 1
	for i := 0; i < len(workflowModels); i++ {
		workflowModels[i].WorkflowVersion = version
	}
	if err := repo.db.Model(&WorkflowModel{}).Create(workflowModels).Error; err != nil {
		return err
	}
	return nil
}

func (repo *grpcRepo) CreateWorkflowInstance(workflowKey string, workflowVersion int, workflowVariables []WorkflowVariable) (uint, bool, string) {
	workflowInstance := &WorkflowInstance{}
	workflowInstance.WorkflowKey = workflowKey
	if workflowVersion != 0 {
		workflowInstance.WorkflowVersion = workflowVersion
	} else {
		workflowModel := &WorkflowModel{}
		if err := repo.db.Table("workflow_models as w").Select("w.workflow_version").Last(workflowModel, "w.workflow_key == ?", workflowKey).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return uint(0), false, err.Error()
			}
		}
		workflowInstance.WorkflowVersion = workflowModel.WorkflowVersion
	}
	if err := repo.db.Create(&workflowInstance).Error; err != nil {
		return uint(0), false, err.Error()
	}
	for _, workflowVariable := range workflowVariables {
		workflowVariable.WorkflowInstanceID = workflowInstance.ID
	}
	if err := repo.db.Model(&WorkflowVariable{}).Create(workflowVariables).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowInstance.ID, true, ""
}
