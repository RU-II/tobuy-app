package router

import (
	"tobuy-app/server/api/controllers"

	"github.com/gorilla/mux"
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
