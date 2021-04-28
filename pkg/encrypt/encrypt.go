package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AESCipher struct {
	block cipher.Block
}

func NewAESCipher(key []byte) (*AESCipher, error) {
	if len(key) != 32 {
		return nil, errors.New("Cipher key size should be 32")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &AESCipher{block}, nil
}

func (a *AESCipher) Encrypt(plaintext string) (string, error) {
	encryptByteArray := make([]byte, aes.BlockSize+len(plaintext))

	iv := encryptByteArray[:aes.BlockSize]
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(a.block, iv)
	stream.XORKeyStream(encryptByteArray[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(encryptByteArray), nil
}

func (a *AESCipher) Decrypt(encrypted string) (string, error) {
	b, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	byteString := []byte(b)

	decryptByteArray := make([]byte, len(byteString))
	iv := byteString[:aes.BlockSize]

	stream := cipher.NewCFBDecrypter(a.block, iv)
	stream.XORKeyStream(decryptByteArray, byteString[aes.BlockSize:])
	return string(bytes.Trim(decryptByteArray, "\x00")), nil
}
