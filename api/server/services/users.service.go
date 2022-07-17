package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"tobuy-app/api/server/models"
	"tobuy-app/api/server/repositories"
	"tobuy-app/api/server/utils/logic"
)

type UsersService interface {
	// ユーザーに関する処理
	UpdateUser(w http.ResponseWriter, r *http.Request, userId int) (models.BaseUserResponse, error)
	UpdatePassword(w http.ResponseWriter, r *http.Request, userId int) (models.BaseUserResponse, error)
	DeleteUser(w http.ResponseWriter, r *http.Request, userId int) error

	// レスポンス送信
	SendUserResponse(w http.ResponseWriter, user *models.BaseUserResponse)
	SendDeleteUserResponse(w http.ResponseWriter)
}

type usersService struct {
	ur repositories.UserRepository
	ul logic.UsersLogic
	rl logic.ResponseLogic
	al logic.AuthLogic
}

func NewUsersService(ur repositories.UserRepository, ul logic.UsersLogic, rl logic.ResponseLogic, al logic.AuthLogic) UsersService {
	return &usersService{ur, ul, rl, al}
}

func (us *usersService) UpdateUser(w http.ResponseWriter, r *http.Request, userId int) (models.BaseUserResponse, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの読み取り処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}
	var updateUserRequest models.UpdateUserRequest
	if err := json.Unmarshal(reqBody, &updateUserRequest); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの構造体への変換処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	var updatedUser models.User
	updatedUser.Name = updateUserRequest.Name
	updatedUser.Email = updateUserRequest.Email

	if err := us.ur.UpdateUser(&updatedUser, userId); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "ユーザーの更新に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	var user models.User
	if err := us.ur.GetUserById(&user, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "ユーザー取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	userResponse := us.ul.CreateUserResponse(&user)

	return userResponse, nil
}

func (us *usersService) UpdatePassword(w http.ResponseWriter, r *http.Request, userId int) (models.BaseUserResponse, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの読み取り処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}
	var updatePasswordRequest models.UpdatePasswordRequest
	if err := json.Unmarshal(reqBody, &updatePasswordRequest); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの構造体への変換処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	// ユーザー認証
	var user models.User
	if err := us.ur.GetUserById(&user, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "ユーザー取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePasswordRequest.Password)); err != nil {
		statusCode := http.StatusUnauthorized
		errMessage := "パスワードが間違っています。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	var updatedUser models.User
	updatedUser.Password = string(us.al.CreateHashPassword(updatePasswordRequest.NewPassword))
	// 空指定のパラメータは更新されない
	if err := us.ur.UpdateUser(&updatedUser, userId); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "ユーザーの更新に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseUserResponse{}, err
	}

	userResponse := us.ul.CreateUserResponse(&user)

	return userResponse, nil
}

func (us *usersService) DeleteUser(w http.ResponseWriter, r *http.Request, userId int) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの読み取り処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}
	var deleteUserRequest models.DeleteUserRequest
	if err := json.Unmarshal(reqBody, &deleteUserRequest); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの構造体への変換処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}

	// ユーザー認証
	var user models.User
	if err := us.ur.GetUserById(&user, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "ユーザー取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(deleteUserRequest.Password)); err != nil {
		statusCode := http.StatusUnauthorized
		errMessage := "パスワードが間違っています。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}

	if err := us.ur.DeleteUser(userId); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "ユーザーの削除に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		us.rl.SendResponse(w, us.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}
	return nil
}

func (us *usersService) SendUserResponse(w http.ResponseWriter, user *models.BaseUserResponse) {
	var response models.UserResponse
	response.User = *user
	responseBody, _ := json.Marshal(response)
	us.rl.SendResponse(w, responseBody, http.StatusOK)
}

func (us *usersService) SendDeleteUserResponse(w http.ResponseWriter) {
	us.rl.SendNotBodyResponse(w)
}
