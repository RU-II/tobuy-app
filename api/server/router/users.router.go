package router

import (
	"tobuy-app/api/server/controllers"

	"github.com/gorilla/mux"
)

type UsersRouter interface {
	SetUsersRouting(router *mux.Router)
}

type usersRouter struct {
	usc controllers.UsersController
}

func NewUsersRouter(usc controllers.UsersController) UsersRouter {
	return &usersRouter{usc}
}

func (usr *usersRouter) SetUsersRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/users", usr.usc.UsersPage).Methods("GET")
}
