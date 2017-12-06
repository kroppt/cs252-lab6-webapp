package http

import (
	"log"
	"net/http"
)

// StartServer will start the server using the given port string.
// Example: http.StartServer(":8080")
func StartServer(p string) {
	r := newRouter()
	checkError(http.ListenAndServe(p, r))
}

func checkError(err error) {
	log.Fatal(err)
}
