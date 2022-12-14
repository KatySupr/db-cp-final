package auth

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, email, password, role string) error
	SignIn(ctx context.Context, email, password, role string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
