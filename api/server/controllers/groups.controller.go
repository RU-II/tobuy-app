package controllers

import (
	"fmt"
	"net/http"
)

type IGroupsController interface {
	GroupsPage(w http.ResponseWriter, r *http.Request)
}

type GroupsController struct{}

func NewGroupsController() *GroupsController {
	return &GroupsController{}
}

func (gc *GroupsController) GroupsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Groups")
	fmt.Println("Groups endpoint is hooked!")
}
