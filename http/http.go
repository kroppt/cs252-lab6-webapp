package http

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

// StartServer will start the server using the given port string.
// Example: http.StartServer(":8080")
func StartServer(p string) {
	r := cors.Default().Handler(newRouter())
	checkError(http.ListenAndServe(p, r))
}

func checkError(err error) {
	log.Fatal(err)
}
