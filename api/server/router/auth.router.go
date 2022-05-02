package router

import (
	"tobuy-app/api/server/controllers"

	"github.com/gorilla/mux"
)

type AuthRouter interface {
	SetAuthRouting(router *mux.Router)
}

type authRouter struct {
	auc controllers.AuthController
}

func NewAuthRouter(auc controllers.AuthController) AuthRouter {
	return &authRouter{auc}
}

func (aur *authRouter) SetAuthRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/auth", aur.auc.AuthPage).Methods("GET")
}
