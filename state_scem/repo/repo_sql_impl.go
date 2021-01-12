package repo

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
)

var errSQLRepo = errors.New("Unable to handle SQL Repo Request")

type sqlRepo struct {
	db     *gorm.DB
	logger log.Logger
}

// NewSQLRepo function
func NewSQLRepo(db *gorm.DB, logger log.Logger) Repository {
	return &sqlRepo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

// ------------------------------------ WorkflowModel ------------------------------------

func (repo *sqlRepo) CreateWorkflowModel(_ context.Context, workflowModel WorkflowModel) (uint, bool, error) {
	workflowModel.ID = 0
	if err := repo.db.Create(&workflowModel).Error; err != nil {
		return uint(0), false, err
	}
	return workflowModel.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowModelLastestVersionByProcessID(_ context.Context, processID string) (WorkflowModel, bool, error) {
	workflowModel := WorkflowModel{}
	if err := repo.db.Table("workflow_models").Last(&workflowModel, "workflow_process_id LIKE ?", processID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return WorkflowModel{}, true, nil
		}
		return WorkflowModel{}, false, err
	}
	return workflowModel, true, nil
}

//  ------------------------------------WorkflowInstance ------------------------------------

func (repo *sqlRepo) CreateWorkflowInstance(_ context.Context, workflowInstance WorkflowInstance) (uint, bool, error) {
	if err := repo.db.Create(&workflowInstance).Error; err != nil {
		return uint(0), false, err
	}
	return workflowInstance.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowInstanceList(_ context.Context) ([]WorkflowInstance, bool, error) {
	workflowInstance := []WorkflowInstance{}
	if err := repo.db.Order("id asc").Find(&workflowInstance).Error; err != nil {
		return nil, false, err
	}
	return workflowInstance, true, nil
}

func (repo *sqlRepo) GetWorkflowInstance(_ context.Context, id uint) (WorkflowInstance, bool, error) {
	workflowInstance := WorkflowInstance{}
	if err := repo.db.First(&workflowInstance, id).Error; err != nil {
		return WorkflowInstance{}, false, err
	}
	return workflowInstance, true, nil
}

func (repo *sqlRepo) UpdateWorkflowInstance(_ context.Context, id uint, workflowInstance WorkflowInstance) (bool, error) {
	workflowInstance.ID = id
	if err := repo.db.Model(&workflowInstance).Updates(workflowInstance).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ------------------------------------ WorkflowRuningPath ------------------------------------
func (repo *sqlRepo) CreateWorkflowRuningPath(_ context.Context, workflowRuningPath WorkflowRuningPath) (uint, bool, error) {
	workflowRuningPath.ID = 0
	if err := repo.db.Create(&workflowRuningPath).Error; err != nil {
		return uint(0), false, err
	}
	return workflowRuningPath.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowRuningPathList(_ context.Context) ([]WorkflowRuningPath, bool, error) {
	workflowRuningPath := []WorkflowRuningPath{}
	if err := repo.db.Order("id asc").Find(&workflowRuningPath).Error; err != nil {
		return nil, false, err
	}
	return workflowRuningPath, true, nil
}

func (repo *sqlRepo) GetWorkflowRuningPathListByWFInstanceID(_ context.Context, workflowInstanceID uint) ([]WorkflowRuningPath, bool, error) {
	workflowRuningPath := []WorkflowRuningPath{}
	if err := repo.db.Order("id asc").Find(&workflowRuningPath, "workflow_instance_id = ?", workflowInstanceID).Error; err != nil {
		return nil, false, err
	}
	return workflowRuningPath, true, nil
}

func (repo *sqlRepo) UpdateWorkflowRuningPath(_ context.Context, id uint, workflowRuningPath WorkflowRuningPath) (bool, error) {
	workflowRuningPath.ID = id
	if err := repo.db.Model(&workflowRuningPath).Updates(workflowRuningPath).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ------------------------------------ WorkflowVariable ------------------------------------

func (repo *sqlRepo) CreateWorkflowVariable(_ context.Context, workflowVariable WorkflowVariable) (uint, bool, error) {
	workflowVariable.ID = 0
	if err := repo.db.Create(&workflowVariable).Error; err != nil {
		return uint(0), false, err
	}
	return workflowVariable.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowVariableList(_ context.Context) ([]WorkflowVariable, bool, error) {
	workflowVariable := []WorkflowVariable{}
	if err := repo.db.Order("id asc").Find(&workflowVariable).Error; err != nil {
		return nil, false, err
	}
	return workflowVariable, true, nil
}

func (repo *sqlRepo) GetWorkflowVariableListByWFInstanceID(_ context.Context, workflowInstanceID uint) ([]WorkflowVariable, bool, error) {
	workflowVariable := []WorkflowVariable{}
	if err := repo.db.Order("id asc").Find(&workflowVariable, "workflow_instance_id = ?", workflowInstanceID).Error; err != nil {
		return nil, false, err
	}
	return workflowVariable, true, nil
}

func (repo *sqlRepo) UpdateWorkflowVariable(_ context.Context, id uint, workflowVariable WorkflowVariable) (bool, error) {
	workflowVariable.ID = id
	if err := repo.db.Model(&workflowVariable).Updates(workflowVariable).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ------------------------------------ WorkflowJobQueue ------------------------------------

func (repo *sqlRepo) CreateWorkflowJobQueue(_ context.Context, workflowJobQueue WorkflowJobQueue) (uint, bool, error) {
	workflowJobQueue.ID = 0
	if err := repo.db.Create(&workflowJobQueue).Error; err != nil {
		return uint(0), false, err
	}
	return workflowJobQueue.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowJobQueueList(_ context.Context) ([]WorkflowJobQueue, bool, error) {
	workflowJobQueue := []WorkflowJobQueue{}
	if err := repo.db.Order("id asc").Find(&workflowJobQueue).Error; err != nil {
		return nil, false, err
	}
	return workflowJobQueue, true, nil
}

func (repo *sqlRepo) GetWorkflowJobQueueListByMappingName(_ context.Context, mappingName string) ([]WorkflowJobQueue, bool, error) {
	workflowJobQueue := []WorkflowJobQueue{}
	if err := repo.db.Order("id asc").Find(&workflowJobQueue, "mapping_name = ?", mappingName).Error; err != nil {
		return nil, false, err
	}
	return workflowJobQueue, true, nil
}

func (repo *sqlRepo) UpdateWorkflowJobQueue(_ context.Context, id uint, workflowJobQueue WorkflowJobQueue) (bool, error) {
	workflowJobQueue.ID = id
	if err := repo.db.Model(&workflowJobQueue).Updates(workflowJobQueue).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ------------------------------------ WorkflowMessageQueue ------------------------------------

func (repo *sqlRepo) CreateWorkflowMessageQueue(_ context.Context, workflowMessageQueue WorkflowMessageQueue) (uint, bool, error) {
	workflowMessageQueue.ID = 0
	if err := repo.db.Create(&workflowMessageQueue).Error; err != nil {
		return uint(0), false, err
	}
	return workflowMessageQueue.ID, true, nil
}

func (repo *sqlRepo) GetWorkflowMessageQueueList(_ context.Context) ([]WorkflowMessageQueue, bool, error) {
	workflowMessageQueue := []WorkflowMessageQueue{}
	if err := repo.db.Order("id asc").Find(&workflowMessageQueue).Error; err != nil {
		return nil, false, err
	}
	return workflowMessageQueue, true, nil
}

func (repo *sqlRepo) GetWorkflowMessageQueueListByMappingName(_ context.Context, mappingName string) ([]WorkflowMessageQueue, bool, error) {
	workflowMessageQueue := []WorkflowMessageQueue{}
	if err := repo.db.Order("id asc").Find(&workflowMessageQueue, "mapping_name = ?", mappingName).Error; err != nil {
		return nil, false, err
	}
	return workflowMessageQueue, true, nil
}

func (repo *sqlRepo) UpdateWorkflowMessageQueue(_ context.Context, id uint, workflowMessageQueue WorkflowMessageQueue) (bool, error) {
	workflowMessageQueue.ID = id
	if err := repo.db.Model(&workflowMessageQueue).Updates(workflowMessageQueue).Error; err != nil {
		return false, err
	}
	return true, nil
}
