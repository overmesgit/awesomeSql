package user_grpc

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/overmesgit/awesomeSql/user_service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct {
	storage user_service.Storage
	UnimplementedUserSignUpServer
}

func (s *server) SignUp(ctx context.Context, in *SignUpRequest) (*LoginResponse, error) {
	userObj, err := user_service.SingUp(s.storage, in.Username, in.Password, in.Email, in.Mood)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{UserId: int32(userObj.UserID)}, nil
}

func (s *server) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	userObj, err := user_service.Login(s.storage, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{UserId: int32(userObj.UserID)}, nil
}

func Start() {
	conn := fmt.Sprintf("dbname=gogo user=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
		return
	}
	psqlStorage := user_service.NewPSQLStorage(db)
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterUserSignUpServer(s, &server{storage: psqlStorage})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
