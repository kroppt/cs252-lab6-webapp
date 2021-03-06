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
		"/api/changePassword",
		[]string{"POST"},
		api.Validate(api.ChangePassword),
		"Update existing password",
	},
	route{
		"/api/loginUser",
		[]string{"POST"},
		api.LoginUser,
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
		[]string{"GET"},
		api.GetID,
		"Get ID",
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
