package controllers

import (
	"fmt"
	"net/http"
)

type GroupsController interface {
	GroupsPage(w http.ResponseWriter, r *http.Request)
}

type groupsController struct{}

func NewGroupsController() GroupsController {
	return &groupsController{}
}

func (grc *groupsController) GroupsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Groups")
	fmt.Println("Groups endpoint is hooked!")
}
