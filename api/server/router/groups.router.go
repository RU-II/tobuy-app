package router

import (
	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
)

type IGroupsRouter interface {
	SetGroupsRouting(router *mux.Router)
}

type GroupsRouter struct {
	grc controllers.IGroupsController
}

func NewGroupsRouter(gc controllers.IGroupsController) *GroupsRouter {
	return &GroupsRouter{gc}
}

func (gr *GroupsRouter) SetGroupsRouting(router *mux.Router) {
	router.HandleFunc(basePath+"/groups", gr.grc.GroupsPage).Methods("GET")
}
