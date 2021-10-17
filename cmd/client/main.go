package main

import (
	"context"
	"flag"
	pb "github.com/overmesgit/awesomeSql/user_grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	command  = flag.String("c", "signup", "signup/login flag")
	username = flag.String("u", "", "username")
	password = flag.String("p", "", "password")
	email    = flag.String("e", "", "email")
	mood     = flag.String("m", "", "mood")
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserSignUpClient(conn)

	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if *command == "signup" {
		request := pb.SignUpRequest{
			Username: *username, Password: *password,
			Email: *email, Mood: *mood,
		}
		r, err := c.SignUp(ctx, &request)
		if err != nil {
			log.Fatalf("could not sing up: %v", err)
		}
		log.Printf("%v: %v", *command, r.GetUserId())
	} else {
		request := pb.LoginRequest{
			Email: *email, Password: *password,
		}
		r, err := c.Login(ctx, &request)
		if err != nil {
			log.Fatalf("could not login: %v", err)
		}
		log.Printf("%v: %v", *command, r.GetUserId())
	}

}
