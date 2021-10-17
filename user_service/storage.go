package user_service

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

const (
	EnumMoodHappy   = "happy"
	EnumMoodSad     = "sad"
	EnumMoodNeutral = "neutral"
)

type User struct {
	UserID   int
	Username string
	Email    string
	Mood     string
}

func (u *User) Hash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

var NotFound = errors.New("user not found")

type Storage interface {
	Create(user *User, passwordHash string) (int, error)
	GetUser(userId int) (*User, error)
	CheckPassword(email string, passwordHash string) (*User, error)
}
