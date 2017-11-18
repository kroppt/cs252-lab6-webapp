package http

import (
	"net/http"
)

func StartServer(p string) {
	r := NewRouter()
	http.ListenAndServe(p, r)
}
