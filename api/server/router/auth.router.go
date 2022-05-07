package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type IAuthRouter interface {
	SetAuthRouting(router *mux.Router)
}

type AuthRouter struct {
	ac controllers.IAuthController
}

func NewAuthRouter(ac controllers.IAuthController) *AuthRouter {
	return &AuthRouter{ac}
}

func (ar *AuthRouter) SetAuthRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/signin", ar.ac.SignIn).Methods(http.MethodPost)
	router.HandleFunc(basePath+"/signup", ar.ac.SignUp).Methods(http.MethodPost)
}
