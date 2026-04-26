package security

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 12

// HashPassword hashes a plain-text password using bcrypt.
func HashPassword(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compares a plain-text password against a bcrypt hash.
func CheckPassword(plain, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
