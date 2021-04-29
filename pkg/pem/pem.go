package pem

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateRsaKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	return privateKey, &privateKey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privateKey *rsa.PrivateKey) string {
	privatePem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	return string(privatePem)
}

func ParseRsaPrivateKeyFromPemStr(privatePem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privatePem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	b, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	publicPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: b,
		},
	)

	return string(publicPem), nil
}

func ParseRsaPublicKeyFromPemStr(publicPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}
