package router

import (
	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type IUsersRouter interface {
	SetUsersRouting(router *mux.Router)
}

type UsersRouter struct {
	usc controllers.IUsersController
}

func NewUsersRouter(uc controllers.IUsersController) *UsersRouter {
	return &UsersRouter{uc}
}

func (ur *UsersRouter) SetUsersRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/users", ur.usc.UsersPage).Methods("GET")
}
