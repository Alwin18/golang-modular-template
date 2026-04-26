package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateRandomToken generates a cryptographically secure random hex token.
func GenerateRandomToken(byteLen int) (string, error) {
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// SHA256Hash returns the SHA-256 hash of a string.
func SHA256Hash(input string) string {
	h := sha256.Sum256([]byte(input))
	return hex.EncodeToString(h[:])
}
