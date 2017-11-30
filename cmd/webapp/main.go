// sample code from https://golang.org/doc/articles/wiki/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	http "github.com/kroppt/cs252-lab6-webapp/http"
	webapp "github.com/kroppt/cs252-lab6-webapp/webapp"
	"log"
)

func main() {
	var err error
	db, err := sql.Open("mysql", "server:server@/test")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	webapp.DB = &webapp.DataBase{db}
	http.StartServer(":8080")
}
