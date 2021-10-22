package login_psql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
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

func (s *PSQLStorage) Create(user *login.User, passwordHash string) (int, *login.Error) {
	userObj := &models.User{Username: user.Username, Password: passwordHash, Email: user.Email,
		Mood: null.StringFrom(user.Mood)}
	err := userObj.Insert(context.TODO(), s.db, boil.Infer())
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Code.Name() == "unique_violation" {
				return 0, login.WrapError(err, "user already exists", login.UserAlreadyExistError)
			}
		} else {
			return 0, login.WrapError(err, "internal error", login.InternalError)
		}
	}
	return userObj.UserID, nil
}

func (s *PSQLStorage) GetUser(userId int) (*login.User, *login.Error) {
	userObj, err := models.Users(models.UserWhere.UserID.EQ(userId)).One(context.TODO(), s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, login.WrapError(err, "not found", login.UserNotFoundError)
		} else {
			return nil, login.WrapError(err, "internal error", login.InternalError)
		}
	}
	return &login.User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, nil
}

func (s *PSQLStorage) CheckPassword(email string, passwordHash string) (*login.User, *login.Error) {
	query := models.Users(models.UserWhere.Email.EQ(email), models.UserWhere.Password.EQ(passwordHash))
	userObj, err := query.One(context.TODO(), s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//return nil, login.WrapError(err, "not found", login.UserNotFoundError)
			return nil, login.WrapError(err, "not found", 19)
		} else {
			return nil, login.WrapError(err, "internal error", login.InternalError)
		}
	}
	return &login.User{UserID: userObj.UserID, Username: userObj.Username, Email: userObj.Email,
		Mood: userObj.Mood.String}, nil
}
