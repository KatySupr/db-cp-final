package usecase

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(ctx context.Context, username, password, role string) error {
	args := m.Called(username, password)

	return args.Error(0)
}

func (m *AuthUseCaseMock) SignIn(ctx context.Context, username, password, role string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUseCaseMock) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}
