package controllers

import (
	"fmt"
	"net/http"
)

type UsersController interface {
	UsersPage(w http.ResponseWriter, r *http.Request)
}

type usersController struct{}

func NewUsersController() UsersController {
	return &usersController{}
}

func (usc *usersController) UsersPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users")
	fmt.Println("Users endpoint is hooked!")
}
