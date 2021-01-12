package repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
)

// NewSQLInitRepo function
func NewSQLInitRepo(db *gorm.DB, logger log.Logger) InitRepository {
	return &sqlRepo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

// ------------------------------------ WorkflowModel ------------------------------------

func (repo *sqlRepo) DeleteDatabase(_ context.Context) (err error) {
	return repo.db.Migrator().DropTable(
		&WorkflowModel{},
		&WorkflowInstance{},
		&WorkflowRuningPath{},
		&WorkflowVariable{},
		&WorkflowJobQueue{},
		&WorkflowMessageQueue{},
	)
}

func (repo *sqlRepo) MigrationDatabase(_ context.Context) (err error) {
	return repo.db.AutoMigrate(
		&WorkflowModel{},
		&WorkflowInstance{},
		&WorkflowRuningPath{},
		&WorkflowVariable{},
		&WorkflowJobQueue{},
		&WorkflowMessageQueue{},
	)
}
