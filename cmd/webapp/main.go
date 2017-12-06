// sample code from https://golang.org/doc/articles/wiki/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	http "github.com/kroppt/cs252-lab6-webapp/http"
	webapp "github.com/kroppt/cs252-lab6-webapp/webapp"
	"io/ioutil"
	"log"
)

func main() {
	var err error
	pass, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", "server:"+string(pass)[:len(pass)-1]+"@/test")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	webapp.DB = &webapp.DataBase{db}
	http.StartServer(":8080")
}
