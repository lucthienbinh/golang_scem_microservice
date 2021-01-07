package service

import (
	"context"
)

// Service interface
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
