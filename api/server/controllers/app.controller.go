package controllers

import (
	"fmt"
	"net/http"
)

type IAppController interface {
	RootPage(w http.ResponseWriter, r *http.Request)
}

type AppController struct{}

func NewAppController() *AppController {
	return &AppController{}
}

func (ac *AppController) RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to tobuy-app server")
	fmt.Println("Root endpoint is hooked!")
}
