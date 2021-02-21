package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Client *sql.DB

	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	host     = os.Getenv("MYSQL_HOST")
	port     = os.Getenv("MYSQL_PORT")
	schema   = os.Getenv("MYSQL_SCHEMA")
)

func init() {
	sqlUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host+":"+port,
		schema,
	)
	var connectionError error
	Client, connectionError = sql.Open("mysql", sqlUri)
	if connectionError != nil {
		panic(connectionError)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully connected")
}
