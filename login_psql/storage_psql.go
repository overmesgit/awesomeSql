package login_psql

import (
	"context"
	"database/sql"
	"github.com/overmesgit/awesomeSql/login"
	"github.com/overmesgit/awesomeSql/login_psql/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PSQLStorage struct {
	db *sql.DB
}

func NewPSQLStorage(conn string) (*PSQLStorage, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &PSQLStorage{db}, nil
}

func (s *PSQLStorage) Create(user *login.User, passwordHash string) (int, error) {
	userObj := &models.User{Username: user.Username, Password: passwordHash, Email: user.Email,
		Mood: null.StringFrom(user.Mood)}
	err := userObj.Insert(context.TODO(), s.db, boil.Infer())
	return userObj.UserID, err
}

func (s *PSQLStorage) GetUser(userId int) (*login.User, error) {
	userObj, err := models.Users(models.UserWhere.UserID.EQ(userId)).One(context.TODO(), s.db)
	return &login.User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, err
}

func (s *PSQLStorage) CheckPassword(email string, passwordHash string) (*login.User, error) {
	query := models.Users(models.UserWhere.Email.EQ(email), models.UserWhere.Password.EQ(passwordHash))
	userObj, err := query.One(context.TODO(), s.db)
	if err == sql.ErrNoRows {
		return nil, login.NotFound
	}
	if err != nil {
		return nil, err
	}
	return &login.User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, nil
}
