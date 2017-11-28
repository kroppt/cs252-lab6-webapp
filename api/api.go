package api

import (
	"encoding/json"
	"fmt"
	"log"
	// "github.com/gorilla/mux"
	webapp "github.com/kroppt/cs252-lab6-webapp/webapp"
	"net/http"
)

type Id struct {
	Id string
}

var storedId Id = Id{""}

func GetID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, storedId.Id)
	fmt.Print("GET ")
	fmt.Println(storedId)
	return
}

func PostID(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&storedId)
	if err != nil {
		panic(err)
	}
	fmt.Print("POST ")
	fmt.Println(storedId)
	w.WriteHeader(http.StatusOK)
	return
}

func TestDB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rows, err := webapp.DB.Query("SELECT * FROM list")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Fprintf(w, "list results:\n")
	for rows.Next() {
		var line string
		if err := rows.Scan(&line); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s\n", line)
	}
	return
}
