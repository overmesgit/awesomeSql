package login

import (
	"errors"
	"testing"
)

type TestStorage struct {
}

var db []*User
var passwords []string

func (t *TestStorage) Create(user *User, passwordHash string) (int32, *Error) {
	db = append(db, user)
	passwords = append(passwords, passwordHash)
	return int32(len(db)), nil
}
func (t *TestStorage) GetUser(userId int32) (*User, *Error) {
	for _, u := range db {
		if u.UserID == userId {
			return u, nil
		}
	}
	return nil, WrapError(errors.New(""), "not found", UserNotFoundError)
}

func (t *TestStorage) CheckPassword(email string, passwordHash string) (*User, *Error) {
	for i, u := range db {
		if u.Email == email && passwords[i] == passwordHash {
			return u, nil
		}
	}
	return nil, WrapError(errors.New(""), "not found", UserNotFoundError)
}

func TestUserLogin(t *testing.T) {
	service := UserService{&TestStorage{}}

	password := "asdf"
	req := SignUpRequest{"art", Password(password), "a@a.com", EnumMoodHappy}
	user, err := service.SignUp(req)

	got, err := service.Login(LoginRequest{"a@a.com", ""})
	if err == nil {
		t.Errorf("got %v", err)
	}

	got, err = service.Login(LoginRequest{"a@a.com", Password(password)})
	if err != nil {
		t.Errorf("got %v", err)
	}
	want := user

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUserLoginNotFound(t *testing.T) {
	service := UserService{&TestStorage{}}
	req := SignUpRequest{"art", "asdf", "a@a.com", EnumMoodHappy}
	user, err := service.SignUp(req)

	got, err := service.Login(LoginRequest{"a1@a.com", req.Password})
	var want *User
	if got != want || err.code != UserNotFoundError {
		t.Errorf("got %q want %v", got, want)
	}

	got, err = service.Login(LoginRequest{user.Email, "1234"})
	if got != want || err.code != UserNotFoundError {
		t.Errorf("got %q want %v", got, want)
	}
}
