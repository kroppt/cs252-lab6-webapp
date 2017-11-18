package http

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, r := range routes {
		router.Path(r.Path).Methods(r.Methods...).Handler(r.Handler).Name(r.Name)
	}

	return router
}
