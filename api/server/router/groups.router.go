package router

import (
	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type GroupsRouter interface {
	SetGroupsRouting(router *mux.Router)
}

type groupsRouter struct {
	grc controllers.GroupsController
}

func NewGroupsRouter(grc controllers.GroupsController) GroupsRouter {
	return &groupsRouter{grc}
}

func (grr *groupsRouter) SetGroupsRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/groups", grr.grc.GroupsPage).Methods("GET")
}
