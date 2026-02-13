package crypto

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password using a secure hashing algorithm.
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
