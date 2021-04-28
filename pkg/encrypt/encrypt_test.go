package encrypt

import (
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptAndDecrypt(t *testing.T) {
	cipher, err := NewAESCipher([]byte(strings.Replace(uuid.NewString(), "-", "", -1)))
	assert.NoError(t, err)

	var plaintext string
	faker.FakeData(&plaintext)

	encrypted, err := cipher.Encrypt(plaintext)
	assert.NoError(t, err)
	assert.NotEqual(t, plaintext, encrypted)

	decrypted, err := cipher.Decrypt(encrypted)
	assert.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}
