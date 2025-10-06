package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func SHA256(bytes int) string {
	randomBytes := make([]byte, bytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return ""
	}

	rawCode := base64.RawURLEncoding.EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(rawCode))
	codeHash := fmt.Sprintf("%x", hash[:])

	return codeHash
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || (block.Type != "RSA PRIVATE KEY" && block.Type != "PRIVATE KEY") {
		return nil, fmt.Errorf("failed to decode PEM block containing private key, got type %q", block.Type)
	}

	if block.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	} else if block.Type == "PRIVATE KEY" {
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA private key")
		}
		return rsaKey, nil
	}

	return nil, fmt.Errorf("unsupported key type %q", block.Type)
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || (block.Type != "RSA PUBLIC KEY" && block.Type != "PUBLIC KEY") {
		return nil, fmt.Errorf("failed to decode PEM block containing public key, got type %q", block.Type)
	}

	if block.Type == "RSA PUBLIC KEY" {
		// PKCS#1
		return x509.ParsePKCS1PublicKey(block.Bytes)
	} else if block.Type == "PUBLIC KEY" {
		// PKIX / PKCS#8
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := key.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA public key")
		}
		return rsaKey, nil
	}

	return nil, fmt.Errorf("unsupported public key type %q", block.Type)
}
