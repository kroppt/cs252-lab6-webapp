package http

import (
	"github.com/kroppt/cs252-lab6-webapp/api"
	"net/http"
)

type route struct {
	Path    string
	Methods []string
	Handler http.HandlerFunc
	Name    string
}

var routes = []route{
	route{
		"/",
		[]string{"GET"},
		api.GetID,
		"GetID",
	},
	route{
		"/",
		[]string{"POST"},
		api.PostID,
		"PostID",
	},
	route{
		"/testdb",
		[]string{"GET"},
		api.TestDB,
		"TestDB",
	},
}
