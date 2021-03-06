package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

type service struct {
	repostory Repo.Repository
	logger    log.Logger
}

// NewService function
func NewService(rep Repo.Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) DeployWorkflowService(ctx context.Context, workflowModelList []Repo.WorkflowModel) (string, bool, error) {
	logger := log.With(s.logger, "service", "DeployWorkflow")
	uuid, _ := uuid.NewV4()
	workflowKey := uuid.String()

	processID := workflowModelList[0].WorkflowProcessID
	oldWorkflowModel, ok, err := s.repostory.GetWorkflowModelLastestVersionByProcessID(ctx, processID)
	if ok == false {
		level.Error(logger).Log("err", err)
		return "", false, err
	}
	newWorkflowVersion := oldWorkflowModel.WorkflowVersion + 1
	for _, workflowModel := range workflowModelList {
		workflowModel.WorkflowVersion = newWorkflowVersion
		workflowModel.WorkflowKey = workflowKey
		if _, ok, err := s.repostory.CreateWorkflowModel(ctx, workflowModel); ok == false {
			level.Error(logger).Log("err", err)
			return "", false, err
		}
	}
	level.Info(logger).Log("workflow key", workflowKey)
	return workflowKey, true, nil
}

func (s service) CreateWorkflowInstanceService(ctx context.Context, processID string, workflowVariable Repo.WorkflowVariable) (string, uint, bool, error) {
	logger := log.With(s.logger, "service", "CreateWorkflowInstance")

	workflowModel, ok, err := s.repostory.GetWorkflowModelLastestVersionByProcessID(ctx, processID)
	if ok == false {
		level.Error(logger).Log("err", err)
		return "", uint(0), false, err
	}

	workflowInstance := Repo.WorkflowInstance{}
	workflowInstance.WorkflowProcessID = processID
	workflowInstance.WorkflowVersion = workflowModel.WorkflowVersion // Auto get lastest workflow version
	newInstanceID, ok, err := s.repostory.CreateWorkflowInstance(ctx, workflowInstance)
	if ok == false {
		level.Error(logger).Log("err", err)
		return "", uint(0), false, err
	}

	workflowVariable.WorkflowInstanceID = newInstanceID
	newInstanceID, ok, err = s.repostory.CreateWorkflowVariable(ctx, workflowVariable)
	if ok == false {
		level.Error(logger).Log("err", err)
		return "", uint(0), false, err
	}

	level.Info(logger).Log("workflow key", workflowModel.WorkflowKey, "workflow instance id", newInstanceID)
	return workflowModel.WorkflowKey, newInstanceID, true, nil
}
