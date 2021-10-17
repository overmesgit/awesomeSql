package user_service

import (
	"context"
	"database/sql"
	"github.com/overmesgit/awesomeSql/user_service/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PSQLStorage struct {
	db *sql.DB
}

func NewPSQLStorage(db *sql.DB) *PSQLStorage {
	return &PSQLStorage{db}
}

func (s *PSQLStorage) Create(user *User, passwordHash string) (int, error) {
	userObj := &models.User{Username: user.Username, Password: passwordHash, Email: user.Email,
		Mood: null.StringFrom(user.Mood)}
	err := userObj.Insert(context.TODO(), s.db, boil.Infer())
	return userObj.UserID, err
}

func (s *PSQLStorage) GetUser(userId int) (*User, error) {
	userObj, err := models.Users(models.UserWhere.UserID.EQ(userId)).One(context.TODO(), s.db)
	return &User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, err
}

func (s *PSQLStorage) CheckPassword(email string, passwordHash string) (*User, error) {
	query := models.Users(models.UserWhere.Email.EQ(email), models.UserWhere.Password.EQ(passwordHash))
	userObj, err := query.One(context.TODO(), s.db)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, NotFound
	}
	return &User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, err
}
