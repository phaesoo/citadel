package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/phaesoo/keybox/internal/models"
	"github.com/phaesoo/keybox/pkg/pem"
	"github.com/pkg/errors"
)

type repo interface {
	AuthKey(ctx context.Context, userID, keyID string) (models.AuthKey, error)
	SetAuthKey(ctx context.Context, authKey models.AuthKey) error
	DeleteAuthKey(ctx context.Context, userID, keyID string) error
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterKey(ctx context.Context, userID string) (string, error) {
	privateKey, publicKey := pem.GenerateRsaKeyPair(2048)

	pubicPem, err := pem.ExportRsaPublicKeyAsPemStr(publicKey)
	if err != nil {
		return "", errors.Wrap(err, "Export RSA public key")
	}

	authKey := models.AuthKey{
		KeyID:      uuid.NewString(),
		PublicPem:  pubicPem,
		PrivatePem: pem.ExportRsaPrivateKeyAsPemStr(privateKey),
		UserID:     userID,
	}

	if err := s.repo.SetAuthKey(ctx, authKey); err != nil {
		return "", err
	}

	return authKey.KeyID, nil
}

func (s *Service) DeregisterKey(ctx context.Context, userID, keyID string) error {
	return s.repo.DeleteAuthKey(ctx, userID, keyID)
}
