package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/overmesgit/awesomeSql/user_service/models"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	conn := fmt.Sprintf("dbname=gogo user=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
		return
	}

	//boil.SetDB(db)
	all, err := models.Users().All(context.TODO(), db)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(all[0])
	fmt.Println(all[0].Username)
}
