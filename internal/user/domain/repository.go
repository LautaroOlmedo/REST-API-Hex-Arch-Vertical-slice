package domain

import (
	"context"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=domain --inpackage=true
type Repository interface {
	CreateUser(ctx context.Context, name, email, password string) error
	GetAllUsers(ctx context.Context) (map[int]User, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
}
