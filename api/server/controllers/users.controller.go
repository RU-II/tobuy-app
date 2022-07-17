package controllers

import (
	"fmt"
	"net/http"

	"tobuy-app/api/server/services"
)

type UsersController interface {
	UsersPage(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type usersController struct {
	us services.UsersService
	as services.AuthService
}

func NewUsersController(us services.UsersService, as services.AuthService) UsersController {
	return &usersController{us, as}
}

func (uc *usersController) UsersPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users")
	fmt.Println("Users endpoint is hooked!")
}

// UpdateUser godoc
// @Summary      Update User
// @Description  Update user's information
// @Tags         updateUser
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        requestBody  body      models.UpdateUserRequest  true  "Update User Request"
// @Success      200          {object}  models.UserResponse
// @Failure      400
// @Failure      500
// @Router       /me/update [post]
func (uc *usersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uc.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	userResponse, err := uc.us.UpdateUser(w, r, userId)
	if err != nil {
		return
	}

	uc.us.SendUserResponse(w, &userResponse)
}

// UpdatePassword godoc
// @Summary      Update Password
// @Description  Update user's password
// @Tags         updatePassword
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        requestBody  body  models.UpdatePasswordRequest  true  "Update User's Password Request"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /me/update/password [post]
func (uc *usersController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := uc.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	userResponse, err := uc.us.UpdatePassword(w, r, userId)
	if err != nil {
		return
	}

	uc.us.SendUserResponse(w, &userResponse)
}

// DeleteUser godoc
// @Summary      Delete User
// @Description  Delete User
// @Tags         deleteUser
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        requestBody  body  models.DeleteUserRequest  true  "Delete User Request"
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /me/delete [post]
func (uc *usersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uc.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	if err := uc.us.DeleteUser(w, r, userId); err != nil {
		return
	}

	uc.us.SendDeleteUserResponse(w)
}
