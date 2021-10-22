package login

import (
	"crypto/sha256"
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

type UserService struct {
	Storage
}

type Password string

func (p Password) Hash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(p)))
}

type Storage interface {
	Create(user *User, passwordHash string) (int, *Error)
	GetUser(userId int) (*User, *Error)
	CheckPassword(email string, passwordHash string) (*User, *Error)
}
