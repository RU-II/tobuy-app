package router

import (
	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type IAppRouter interface {
	SetAppRouting(router *mux.Router)
}

type AppRouter struct {
	ac controllers.IAppController
}

func NewAppRouter(ac controllers.IAppController) *AppRouter {
	return &AppRouter{ac}
}

func (apr *AppRouter) SetAppRouting(router *mux.Router) {
	router.HandleFunc(basePath, apr.ac.RootPage).Methods("GET")
}
