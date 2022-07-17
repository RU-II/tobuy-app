package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type UsersRouter interface {
	SetUsersRouting(router *mux.Router)
}

type usersRouter struct {
	uc controllers.UsersController
}

func NewUsersRouter(uc controllers.UsersController) UsersRouter {
	return &usersRouter{uc}
}

func (ur *usersRouter) SetUsersRouting(router *mux.Router) {
	// router.HandleFunc(basePath+"/users", ur.uc.UsersPage).Methods("GET")
	// UpdateUser
	router.HandleFunc(basePath+"/me/update", ur.uc.UpdateUser).Methods(http.MethodPost)
	// UpdatePassword
	router.HandleFunc(basePath+"/me/update/password", ur.uc.UpdatePassword).Methods(http.MethodPost)
	// DeleteUser
	router.HandleFunc(basePath+"/me/delete", ur.uc.DeleteUser).Methods(http.MethodPost)
}
