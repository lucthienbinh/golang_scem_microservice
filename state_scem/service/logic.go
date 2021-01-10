package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
	"gorm.io/gorm"
)

type service struct {
	repostory Repo.GRPCRepository
	logger    log.Logger
}

// NewService function
func NewService(rep Repo.GRPCRepository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) CreateWorkflowModel(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return email, nil
}

func (repo *grpcRepo) CreateWorkflowModel(workflowModels []WorkflowModel) (bool, string) {
	workflowKey := workflowModels[0].WorkflowKey
	workflowModel := &WorkflowModel{}
	version := int(1)
	if err := repo.db.Table("workflow_models as w").Select("w.workflow_version").Last(workflowModel, "w.workflow_key == ?", workflowKey).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err.Error()
		}
	}
	version = workflowModel.WorkflowVersion + 1
	for i := 0; i < len(workflowModels); i++ {
		workflowModels[i].WorkflowVersion = version
	}
	if err := repo.db.Model(&WorkflowModel{}).Create(workflowModels).Error; err != nil {
		return false, err.Error()
	}
	return true, ""
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
