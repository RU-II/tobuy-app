package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type AuthRouter interface {
	SetAuthRouting(router *mux.Router)
}

type authRouter struct {
	ac controllers.AuthController
}

func NewAuthRouter(ac controllers.AuthController) AuthRouter {
	return &authRouter{ac}
}

func (ar *authRouter) SetAuthRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/signin", ar.ac.SignIn).Methods(http.MethodPost)
	router.HandleFunc(basePath+"/signup", ar.ac.SignUp).Methods(http.MethodPost)
}
