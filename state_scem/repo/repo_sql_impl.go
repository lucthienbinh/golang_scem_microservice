package repo

import (
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

func (repo *sqlRepo) CreateWorkflowModel(workflowModel WorkflowModel) (uint, bool, string) {
	workflowModel.ID = 0
	if err := repo.db.Create(&workflowModel).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowModel.ID, true, ""
}

//  ------------------------------------WorkflowInstance ------------------------------------

func (repo *sqlRepo) CreateWorkflowInstance(workflowKey string, workflowVersion int) (uint, bool, string) {
	workflowInstance := &WorkflowInstance{}
	workflowInstance.WorkflowKey = workflowKey
	workflowInstance.WorkflowVersion = workflowVersion
	if err := repo.db.Create(workflowInstance).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowInstance.ID, true, ""
}

func (repo *sqlRepo) GetWorkflowInstanceList() ([]WorkflowInstance, bool, string) {
	workflowInstance := []WorkflowInstance{}
	if err := repo.db.Order("id asc").Find(&workflowInstance).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowInstance, true, ""
}

func (repo *sqlRepo) GetWorkflowInstance(id uint) (WorkflowInstance, bool, string) {
	workflowInstance := WorkflowInstance{}
	if err := repo.db.First(&workflowInstance, id).Error; err != nil {
		return WorkflowInstance{}, false, err.Error()
	}
	return workflowInstance, true, ""
}

func (repo *sqlRepo) UpdateWorkflowInstance(id uint, workflowInstance WorkflowInstance) (bool, string) {
	workflowInstance.ID = id
	if err := repo.db.Model(&workflowInstance).Updates(workflowInstance).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ------------------------------------ WorkflowRuningPath ------------------------------------
func (repo *sqlRepo) CreateWorkflowRuningPath(workflowRuningPath WorkflowRuningPath) (uint, bool, string) {
	workflowRuningPath.ID = 0
	if err := repo.db.Create(&workflowRuningPath).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowRuningPath.ID, true, ""
}

func (repo *sqlRepo) GetWorkflowRuningPathList() ([]WorkflowRuningPath, bool, string) {
	workflowRuningPath := []WorkflowRuningPath{}
	if err := repo.db.Order("id asc").Find(&workflowRuningPath).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowRuningPath, true, ""
}

func (repo *sqlRepo) GetWorkflowRuningPathListByWFInstanceID(workflowInstanceID uint) ([]WorkflowRuningPath, bool, string) {
	workflowRuningPath := []WorkflowRuningPath{}
	if err := repo.db.Order("id asc").Find(&workflowRuningPath, "workflow_instance_id = ?", workflowInstanceID).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowRuningPath, true, ""
}

func (repo *sqlRepo) UpdateWorkflowRuningPath(id uint, workflowRuningPath WorkflowRuningPath) (bool, string) {
	workflowRuningPath.ID = id
	if err := repo.db.Model(&workflowRuningPath).Updates(workflowRuningPath).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ------------------------------------ WorkflowVariable ------------------------------------

func (repo *sqlRepo) CreateWorkflowVariable(workflowVariable WorkflowVariable) (uint, bool, string) {
	workflowVariable.ID = 0
	if err := repo.db.Create(&workflowVariable).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowVariable.ID, true, ""
}

func (repo *sqlRepo) GetWorkflowVariableList() ([]WorkflowVariable, bool, string) {
	workflowVariable := []WorkflowVariable{}
	if err := repo.db.Order("id asc").Find(&workflowVariable).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowVariable, true, ""
}

func (repo *sqlRepo) GetWorkflowVariableListByWFInstanceID(workflowInstanceID uint) ([]WorkflowVariable, bool, string) {
	workflowVariable := []WorkflowVariable{}
	if err := repo.db.Order("id asc").Find(&workflowVariable, "workflow_instance_id = ?", workflowInstanceID).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowVariable, true, ""
}

func (repo *sqlRepo) UpdateWorkflowVariable(id uint, workflowVariable WorkflowVariable) (bool, string) {
	workflowVariable.ID = id
	if err := repo.db.Model(&workflowVariable).Updates(workflowVariable).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ------------------------------------ WorkflowJobQueue ------------------------------------

func (repo *sqlRepo) CreateWorkflowJobQueue(workflowJobQueue WorkflowJobQueue) (uint, bool, string) {
	workflowJobQueue.ID = 0
	if err := repo.db.Create(&workflowJobQueue).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowJobQueue.ID, true, ""
}

func (repo *sqlRepo) GetWorkflowJobQueueList() ([]WorkflowJobQueue, bool, string) {
	workflowJobQueue := []WorkflowJobQueue{}
	if err := repo.db.Order("id asc").Find(&workflowJobQueue).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowJobQueue, true, ""
}

func (repo *sqlRepo) GetWorkflowJobQueueListByName(name string) ([]WorkflowJobQueue, bool, string) {
	workflowJobQueue := []WorkflowJobQueue{}
	if err := repo.db.Order("id asc").Find(&workflowJobQueue, "name = ?", name).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowJobQueue, true, ""
}

func (repo *sqlRepo) UpdateWorkflowJobQueue(id uint, workflowJobQueue WorkflowJobQueue) (bool, string) {
	workflowJobQueue.ID = id
	if err := repo.db.Model(&workflowJobQueue).Updates(workflowJobQueue).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ------------------------------------ WorkflowMessageQueue ------------------------------------

func (repo *sqlRepo) CreateWorkflowMessageQueue(workflowMessageQueue WorkflowMessageQueue) (uint, bool, string) {
	workflowMessageQueue.ID = 0
	if err := repo.db.Create(&workflowMessageQueue).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowMessageQueue.ID, true, ""
}

func (repo *sqlRepo) GetWorkflowMessageQueueList() ([]WorkflowMessageQueue, bool, string) {
	workflowMessageQueue := []WorkflowMessageQueue{}
	if err := repo.db.Order("id asc").Find(&workflowMessageQueue).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowMessageQueue, true, ""
}

func (repo *sqlRepo) GetWorkflowMessageQueueListbyName(name string) ([]WorkflowMessageQueue, bool, string) {
	workflowMessageQueue := []WorkflowMessageQueue{}
	if err := repo.db.Order("id asc").Find(&workflowMessageQueue, "name = ?", name).Error; err != nil {
		return nil, false, err.Error()
	}
	return workflowMessageQueue, true, ""
}

func (repo *sqlRepo) UpdateWorkflowMessageQueue(id uint, workflowMessageQueue WorkflowMessageQueue) (bool, string) {
	workflowMessageQueue.ID = id
	if err := repo.db.Model(&workflowMessageQueue).Updates(workflowMessageQueue).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
}
