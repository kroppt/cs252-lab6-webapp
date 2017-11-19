package http

import (
	"net/http"
)

func StartServer(p string) {
	r := newRouter()
	http.ListenAndServe(p, r)
}
