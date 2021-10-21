package login

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	Create(user *User, passwordHash string) (int, error)
	GetUser(userId int) (*User, error)
	CheckPassword(email string, passwordHash string) (*User, error)
}

func GetFormErrors(err error) map[string][]map[string]string {
	errorResp := map[string][]map[string]string{}
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errorResp[fieldErr.Field()] = append(errorResp[fieldErr.StructField()],
			map[string]string{
				"field":   fieldErr.Field(),
				"tag":     fieldErr.Tag(),
				"param":   fieldErr.Param(),
				"message": fieldErr.Error(),
			})
		fmt.Println(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
	}
	return errorResp
}
