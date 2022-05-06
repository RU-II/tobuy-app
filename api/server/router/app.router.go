package router

import (
	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type AppRouter interface {
	SetAppRouting(router *mux.Router)
}

type appRouter struct {
	apc controllers.AppController
}

func NewAppRouter(apc controllers.AppController) AppRouter {
	return &appRouter{apc}
}

func (apr *appRouter) SetAppRouting(router *mux.Router) {
	router.HandleFunc(basePath, apr.apc.RootPage).Methods("GET")
}
