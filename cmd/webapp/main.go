// sample code from https://golang.org/doc/articles/wiki/

package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	http "github.com/kroppt/cs252-lab6-webapp/http"
	webapp "github.com/kroppt/cs252-lab6-webapp/webapp"
)

func main() {
	pass, err := ioutil.ReadFile("pass.txt")
	checkError(err)
	// remove newline from password
	db, err := sql.Open("mysql", "server:"+string(pass)[:len(pass)-1]+"@/lab6_db?parseTime=true")
	checkError(err)
	err = db.Ping()
	checkError(err)
	webapp.DataBase.DB = db
	http.StartServer(":8080")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
