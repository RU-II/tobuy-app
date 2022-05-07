package controllers

import (
	"fmt"
	"net/http"
)

type IUsersController interface {
	UsersPage(w http.ResponseWriter, r *http.Request)
}

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (uc *UsersController) UsersPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users")
	fmt.Println("Users endpoint is hooked!")
}
