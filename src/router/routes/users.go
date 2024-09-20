package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoute = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Function:  controllers.CreateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodGet,
		Function:  controllers.SearchUsers,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		Function:  controllers.SearchOnlyOneUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodPut,
		Function:  controllers.UpdateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		NeedsAuth: false,
	},
}
