package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
