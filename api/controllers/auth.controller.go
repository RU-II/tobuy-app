package controllers

import (
	"fmt"
	"net/http"
)

type AuthController interface {
	AuthPage(w http.ResponseWriter, r *http.Request)
}

type authController struct{}

func NewAuthController() AuthController {
	return &authController{}
}

func (auc *authController) AuthPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Auth")
	fmt.Println("Auth endpoint is hooked!")
}
