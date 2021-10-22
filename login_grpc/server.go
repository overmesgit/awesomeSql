package login_grpc

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/overmesgit/awesomeSql/login"
	"github.com/overmesgit/awesomeSql/login_psql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

type server struct {
	service login.UserService
	UnimplementedUserSignUpServer
}

func (s *server) SignUp(ctx context.Context, in *SignUpRequest) (*LoginResponse, error) {
	req := login.SignUpRequest{
		Username: in.Username, Password: login.Password(in.Password),
		Email: in.Email, Mood: in.Mood}
	userObj, err := s.service.SingUp(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &LoginResponse{UserId: int32(userObj.UserID)}, nil
}

func (s *server) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	userObj, err := s.service.Login(login.LoginRequest{
		Email: in.Email, Password: login.Password(in.Password),
	})
	if err != nil {
		switch err.Code() {
		case login.UserNotFoundError:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, err
		}
	}
	return &LoginResponse{UserId: int32(userObj.UserID)}, nil
}

func Start() {
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_BASE"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	psqlStorage, err := login_psql.NewPSQLStorage(conn)
	if err != nil {
		log.Fatal(err)
		return
	}
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterUserSignUpServer(s, &server{service: login.UserService{Storage: psqlStorage}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
