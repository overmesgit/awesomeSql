package user_service

import (
	"crypto/sha256"
	"fmt"
)
import _ "crypto/sha256"

func Login(storage Storage, email string, password string) (user *User, err error) {
	passwordHash := sha256.Sum256([]byte(password))
	return storage.CheckPassword(email, fmt.Sprintf("%x", passwordHash))
}
