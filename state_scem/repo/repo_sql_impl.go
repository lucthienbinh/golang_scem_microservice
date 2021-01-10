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

func (repo *sqlRepo) CreateWorkflowModel(workflowModel WorkflowModel) (uint, bool, string) {
	workflowModel.ID = 0
	if err := repo.db.Create(&workflowModel).Error; err != nil {
		return uint(0), false, err.Error()
	}
	return workflowModel.ID, true, ""
}

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

func (repo *sqlRepo) GetWorkflowRuningPath(id uint) (WorkflowRuningPath, bool, string) {
	workflowRuningPath := WorkflowRuningPath{}
	if err := repo.db.First(&workflowRuningPath, id).Error; err != nil {
		return WorkflowRuningPath{}, false, err.Error()
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
