package password

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

const (
	SaltLength = 32
)

func GenerateSalt() (string, error) {
	salt := make([]byte, SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.New("failed to generate salt")
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	combination := password + salt
	hasher := sha256.New()
	hasher.Write([]byte(combination))
	return hex.EncodeToString(hasher.Sum(nil))
}

func VerifyPassword(password, salt, storedHash string) bool {
	expectedHash := HashPassword(password, salt)
	return expectedHash == storedHash
}

func GeneratePasswordAndSalt(password string) (hashedPassword string, salt string, err error) {
	salt, err = GenerateSalt()
	if err != nil {
		return "", "", err
	}
	
	hashedPassword = HashPassword(password, salt)
	return hashedPassword, salt, nil
}
