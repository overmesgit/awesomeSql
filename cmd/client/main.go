package main

import (
	"context"
	"flag"
	pb "github.com/overmesgit/awesomeSql/login_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"os"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	target := os.Getenv("HOST")
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserSignUpClient(conn)

	flag.Parse()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
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
		log.Printf("%v: %v", *command, r.GetUser())
	} else {
		request := pb.LoginRequest{
			Email: *email, Password: *password,
		}
		r, err := c.Login(ctx, &request)

		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				log.Fatalf("could not login: c: %v m: %v d: %v", st.Code(), st.Message(), st.Details())
			} else {
				log.Fatalf("could not login: %v", err)
			}
		}
		log.Printf("%v: %v", *command, r.GetUser())
	}

}
