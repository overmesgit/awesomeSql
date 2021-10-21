package login

import (
	"testing"
)

type TestStorage struct {
}

var db []*User
var passwords []string

func (t *TestStorage) Create(user *User, passwordHash string) (int, error) {
	db = append(db, user)
	passwords = append(passwords, passwordHash)
	return len(db), nil
}
func (t *TestStorage) GetUser(userId int) (*User, error) {
	for _, u := range db {
		if u.UserID == userId {
			return u, nil
		}
	}
	return nil, NotFound
}

func (t *TestStorage) CheckPassword(email string, passwordHash string) (*User, error) {
	for i, u := range db {
		if u.Email == email && passwords[i] == passwordHash {
			return u, nil
		}
	}
	return nil, NotFound
}

func TestUserLogin(t *testing.T) {
	service := UserService{&TestStorage{}}

	password := "asdf"
	req := SignUpRequest{"art", Password(password), "a@a.com", EnumMoodHappy}
	user, err := service.SingUp(req)

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
	user, err := service.SingUp(req)

	got, err := service.Login(LoginRequest{"a1@a.com", "asdf"})
	var want *User
	wantError := NotFound
	if got != want && err != wantError {
		t.Errorf("got %q want %q", got, want)
	}

	got, err = service.Login(LoginRequest{user.Email, "1234"})
	if got != want && err != wantError {
		t.Errorf("got %q want %q", got, want)
	}
}
