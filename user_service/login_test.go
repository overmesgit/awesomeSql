package user_service

import (
	"testing"
)

type TestStorage struct {
}

var db []*User

func (t *TestStorage) Create(user *User) (int, error) {
	db = append(db, user)
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
	for _, u := range db {
		if u.Email == email && u.Hash() == passwordHash {
			return u, nil
		}
	}
	return nil, NotFound
}

func TestUserLogin(t *testing.T) {
	storage := &TestStorage{}

	user, err := SingUp(storage, "art", "asdf", "a@a.com", EnumMoodHappy)

	got, err := Login(storage, "a@a.com", "asdf")
	if err != nil {
		t.Errorf("got %q", got)
	}
	want := user

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUserLoginNotFound(t *testing.T) {
	storage := &TestStorage{}
	user, err := SingUp(storage, "art", "asdf", "a@a.com", EnumMoodHappy)

	got, err := Login(storage, "a1@a.com", user.Password)
	var want *User
	wantError := NotFound
	if got != want && err != wantError {
		t.Errorf("got %q want %q", got, want)
	}

	got, err = Login(storage, user.Email, "1234")
	if got != want && err != wantError {
		t.Errorf("got %q want %q", got, want)
	}
}
