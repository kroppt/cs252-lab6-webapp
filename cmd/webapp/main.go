// sample code from https://golang.org/doc/articles/wiki/

package main

import (
	http "github.com/kroppt/cs252-lab6-webapp/http"
)

func main() {
	http.StartServer(":8080")
}
