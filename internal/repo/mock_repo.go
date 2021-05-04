package repo

import (
	"context"

	"github.com/phaesoo/keybox/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func NewMockRepo() *MockRepo {
	return &MockRepo{}
}

func (m *MockRepo) AuthKey(ctx context.Context, userID, keyID string) (models.AuthKey, error) {
	args := m.MethodCalled("AuthKey", ctx, userID, keyID)
	return args.Get(0).(models.AuthKey), args.Error(1)
}

func (m *MockRepo) SetAuthKey(ctx context.Context, authKey models.AuthKey) error {
	args := m.MethodCalled("SetAuthKey", ctx, authKey)
	return args.Error(0)
}

func (m *MockRepo) DeleteAuthKey(ctx context.Context, userID, keyID string) error {
	args := m.MethodCalled("DeleteAuthKey", ctx, userID, keyID)
	return args.Error(0)
}
