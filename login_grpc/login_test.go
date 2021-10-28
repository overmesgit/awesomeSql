package login_grpc

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/overmesgit/awesomeSql/login"
	"github.com/overmesgit/awesomeSql/login_psql"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

var psqlStorage *login_psql.PSQLStorage

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Second * 5
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "alpine", []string{"POSTGRES_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	if err := pool.Retry(func() error {
		port := resource.GetPort("5432/tcp")
		fmt.Println(port)
		psqlStorage, err = login_psql.NewPSQLStorage("user=postgres password=secret sslmode=disable host=0.0.0.0 port=" + port)
		if err != nil {
			log.Errorf("Can not connect to postgres: %s", err)
			return err
		}
		err = psqlStorage.GetDB().Ping()
		if err != nil {
			log.Errorf("Can not connect to postgres: %s", err)
		}
		return err
	}); err != nil {
		if err := pool.Purge(resource); err != nil {
			log.Errorf("Could not purge resource: %s", err)
		}
		log.Fatalf("Could not connect to database: %s", err)
	}
	driver, err := postgres.WithInstance(psqlStorage.GetDB(), &postgres.Config{})
	if err != nil {
		log.Fatalf("Migration failed: %s", err)
	}
	mig, err := migrate.NewWithDatabaseInstance(
		"file://../login_psql/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migration failed: %s", err)
	}
	err = mig.Up()
	if err != nil {
		log.Fatalf("Migration failed: %s", err)
	}

	code := m.Run()
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestUserLogin(t *testing.T) {
	service := login.UserService{Storage: psqlStorage}

	password := "asdf"
	req := login.SignUpRequest{"art", login.Password(password),
		"a@a.com", login.EnumMoodHappy}
	user, err := service.SignUp(req)

	got, err := service.Login(login.LoginRequest{"a@a.com", ""})
	if err == nil {
		t.Errorf("got %v", err)
	}

	got, err = service.Login(login.LoginRequest{"a@a.com", login.Password(password)})
	if err != nil {
		t.Errorf("got %v", err)
	}
	want := user

	if *got != *want {
		t.Errorf("got %q want %q", got, want)
	}
}

//
//func TestUserLoginNotFound(t *testing.T) {
//	service := UserService{&TestStorage{}}
//	req := SignUpRequest{"art", "asdf", "a@a.com", EnumMoodHappy}
//	user, err := service.SignUp(req)
//
//	got, err := service.Login(LoginRequest{"a1@a.com", "asdf"})
//	var want *User
//	wantError := NotFound
//	if got != want && err != wantError {
//		t.Errorf("got %q want %q", got, want)
//	}
//
//	got, err = service.Login(LoginRequest{user.Email, "1234"})
//	if got != want && err != wantError {
//		t.Errorf("got %q want %q", got, want)
//	}
//}
