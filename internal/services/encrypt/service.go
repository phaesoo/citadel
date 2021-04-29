package encrypt

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	"github.com/phaesoo/keybox/internal/models"
	"github.com/phaesoo/keybox/pkg/pem"
)

type repo interface {
	AuthKey(ctx context.Context, userID, keyID string) (models.AuthKey, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Encrypt(ctx context.Context, userID, keyID string, plaintexts []string) ([]string, error) {
	authKey, err := s.repo.AuthKey(ctx, userID, keyID)
	if err != nil {
		return nil, err
	}

	publicKey, err := pem.ParseRsaPublicKeyFromPemStr(authKey.PublicKey)
	if err != nil {
		return nil, err
	}

	ciphertexts := make([]string, len(plaintexts))

	for i, plaintext := range plaintexts {
		ciphertext, err := rsa.EncryptPKCS1v15(
			rand.Reader,
			publicKey,
			[]byte(plaintext),
		)
		if err != nil {
			return nil, err
		}
		ciphertexts[i] = string(ciphertext)
	}

	return ciphertexts, nil
}
