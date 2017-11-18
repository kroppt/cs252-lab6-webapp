package api

import (
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
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
