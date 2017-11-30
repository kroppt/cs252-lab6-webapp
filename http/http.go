package http

import (
	"log"
	"net/http"
)

func StartServer(p string) {
	r := newRouter()
	if err := http.ListenAndServe(p, r); err != nil {
		log.Fatal(err)
	}
}
