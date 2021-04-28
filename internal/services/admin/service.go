package admin

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterKey() (string, error) {
	keyID := uuid.NewString()
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	return keyID, nil
}
