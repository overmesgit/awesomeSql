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

const (
	EnumUserTypeCustomer = "customer"
	EnumUserTypeSeller   = "seller"
)

type UserType string

type User struct {
	UserID   int32
	Username string
	Email    string
	Mood     string
	Type     UserType
}

type UserService struct {
	Storage
}

type Password string

func (p Password) Hash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(p)))
}

type Storage interface {
	Create(user *User, passwordHash string) (int32, *Error)
	GetUser(userId int32) (*User, *Error)
	CheckPassword(email string, passwordHash string) (*User, *Error)
}
