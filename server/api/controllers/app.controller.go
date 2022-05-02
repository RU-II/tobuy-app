package controllers

import (
	"fmt"
	"net/http"
)

type AppController interface {
	RootPage(w http.ResponseWriter, r *http.Request)
}

type appController struct{}

func NewAppController() AppController {
	return &appController{}
}

func (apc *appController) RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to tobuy-app server")
	fmt.Println("Root endpoint is hooked!")
}
