package http

import (
	"net/http"

	"github.com/kroppt/cs252-lab6-webapp/api"
)

type route struct {
	Path    string
	Methods []string
	Handler http.HandlerFunc
	Name    string
}

var routes = []route{
	route{
		"/api/authUser",
		[]string{"POST"},
		api.AuthUser,
		"Authenticate User",
	},
	route{
		"/",
		[]string{"GET"},
		api.GetID,
		"Get ID",
	},
	route{
		"/api/loginUser",
		[]string{"POST"},
		api.Validate(api.LoginUser),
		"Login User",
	},
	route{
		"/api/logoutUser",
		[]string{"POST"},
		api.Validate(api.LogoutUser),
		"Logout User",
	},
	route{
		"/api/newUser",
		[]string{"POST"},
		api.NewUser,
		"New User",
	},
	route{
		"/",
		[]string{"POST"},
		api.PostID,
		"Post ID",
	},
	route{
		"/testdb",
		[]string{"GET"},
		api.TestDB,
		"Test database",
	},
}
