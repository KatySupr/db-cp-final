package auth

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, email, password string) (*models.User, error)
}
