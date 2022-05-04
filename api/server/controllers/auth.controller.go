package controllers

import (
	"net/http"

	"tobuy-app/api/server/services"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	as services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
	return &authController{as}
}

// SignIn godoc
// @Summary      SignIn
// @Description  SignIn
// @Tags         signIn
// @Accept       json
// @Produce      json
// @Param        requestBody  body      models.SignInRequest  true  "SignIn Request"
// @Success      200          {object}  models.AuthResponse
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Router       /signin [post]
func (ac *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := ac.as.SignIn(w, r)
	if err != nil {
		return
	}

	ac.as.SendAuthResponse(w, &user, http.StatusOK)
}

// SignUp godoc
// @Summary      SignUp
// @Description  SignUp
// @Tags         signUp
// @Accept       json
// @Produce      json
// @Param        requestBody  body      models.SignUpRequest  true  "SignUp Request"
// @Success      201          {object}  models.AuthResponse
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Router       /signup [post]
func (ac *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	createdUser, err := ac.as.SignUp(w, r)
	if err != nil {
		return
	}

	ac.as.SendAuthResponse(w, &createdUser, http.StatusCreated)
}
