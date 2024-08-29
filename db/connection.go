package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pwd := os.Getenv("MYSQL_PASSWORD")
	usr := os.Getenv("USER")
	db, err := sql.Open("mysql", usr+":"+pwd+"@tcp(localhost:3307)/lealtest")

	if err != nil {
		fmt.Println("error validation sql.open arguments")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	fmt.Println("Successful Connection to Database!")
	//defer db.Close()

	return db
}
