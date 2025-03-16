package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
} 

func HashPassword(password string, salt string) string {
	hashPassword := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hashPassword[:])
}

func VerifyPassword(password string, hashedPassword string, salt string) bool {
	hashPassword := HashPassword(password, salt)
	return hashPassword == hashedPassword
}

func GenerateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(salt)
}

