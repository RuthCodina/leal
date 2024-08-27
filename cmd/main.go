package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leal/cmd/api"
)

const defaultPort = "8000"

func main() {

	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	pwd := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pwd+"@tcp(localhost:3306)/lealtest")

	if err != nil {
		fmt.Println("error validation sql.open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	fmt.Println("Successful Connection to Database!")

	// start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	api.Start(port)

}
