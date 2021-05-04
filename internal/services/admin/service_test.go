package admin

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/phaesoo/keybox/internal/models"
	"github.com/phaesoo/keybox/internal/repo"
	"github.com/phaesoo/keybox/pkg/typing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_RegisterKey(t *testing.T) {
	repo := repo.NewMockRepo()
	service := NewService(repo)

	userID := "user-1"

	repo.On("SetAuthKey", mock.Anything, mock.Anything).Return(
		nil,
	).Once()

	keyID, err := service.RegisterKey(context.Background(), userID)
	assert.NoError(t, err)
	assert.True(t, typing.IsValidUUID(keyID))
}

func TestService_DeregisterKey(t *testing.T) {
	repo := repo.NewMockRepo()
	service := NewService(repo)

	userID := "user-1"
	keyID := "key-1"

	authKey := models.AuthKey{}

	assert.NoError(t, faker.FakeData(&authKey))

	ctx := context.Background()

	repo.On("DeleteAuthKey", ctx, userID, keyID).Return(
		nil,
	).Once()

	err := service.DeregisterKey(ctx, userID, keyID)
	assert.NoError(t, err)
}
