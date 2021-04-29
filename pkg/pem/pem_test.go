package pem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExportAndParsePrivateKey(t *testing.T) {
	privateKey, _ := GenerateRsaKeyPair(2048)

	privatePem := ExportRsaPrivateKeyAsPemStr(privateKey)

	parsed, err := ParseRsaPrivateKeyFromPemStr(privatePem)
	assert.NoError(t, err)
	assert.Equal(t, privateKey, parsed)
}

func Test_ExportAndParsePublicKey(t *testing.T) {
	_, publicKey := GenerateRsaKeyPair(2048)

	publicPem, err := ExportRsaPublicKeyAsPemStr(publicKey)
	assert.NoError(t, err)

	parsed, err := ParseRsaPublicKeyFromPemStr(publicPem)
	assert.NoError(t, err)
	assert.Equal(t, publicKey, parsed)
}
