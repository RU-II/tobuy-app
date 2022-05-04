package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"tobuy-app/api/server/models"
	"tobuy-app/api/server/repositories"
	"tobuy-app/api/server/utils/logic"
)

type AuthService interface {
	GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error)
	SignIn(w http.ResponseWriter, r *http.Request) (models.User, error)
	SignUp(w http.ResponseWriter, r *http.Request) (models.User, error)
	SendAuthResponse(w http.ResponseWriter, user *models.User, code int)
}

type authService struct {
	ur repositories.UserRepository
	al logic.AuthLogic
	rl logic.ResponseLogic
	jl logic.JWTLogic
}

func NewAuthService(ur repositories.UserRepository, al logic.AuthLogic, rl logic.ResponseLogic, jl logic.JWTLogic) AuthService {
	return &authService{ur, al, rl, jl}
}

func (as *authService) GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// トークンからuserIdを取得
	userId, err := as.al.GetUserIdFromToken(r)
	if err != nil {
		statusCode := http.StatusUnauthorized
		errMessage := "認証エラー"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return 0, err
	}

	return userId, nil
}

func (as *authService) SignIn(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// RequestのBodyデータを取得
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signInRequestParam models.SignInRequest
	// Unmarshal: jsonを構造体に変換
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	// TODO バリデーション

	// ユーザー認証
	var user models.User
	// emailに紐づくユーザーをチェック
	if err := as.ur.GetUserByEmail(&user, signInRequestParam.Email); err != nil {
		statusCode := http.StatusNotFound
		errMessage := "メールアドレスに該当するユーザーが存在しません。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}
	// パスワード照合
	// CompareHashAndPassword
	// 第一引数: hash化したパスワード
	// 第二引数: 文字列のパスワード
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		statusCode := http.StatusUnauthorized
		errMessage := "パスワードが間違っています。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	return user, nil
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var signUpRequest models.SignUpRequest
	if err := json.Unmarshal(reqBody, &signUpRequest); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストパラメータの構造体への変換処理でエラーが発生しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	// TODO バリデーション

	var users []models.User
	if err := as.ur.GetAllUsersByEmail(&users, signUpRequest.Email); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "DBエラーが発生しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	if len(users) > 0 {
		err := errors.Errorf("email: %w のユーザーは既に登録されています。", signUpRequest.Email)
		statusCode := http.StatusInternalServerError
		errMessage := "入力されたメールアドレスは既に登録されています。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	var newUser models.User
	newUser.Name = signUpRequest.Name
	newUser.Email = signUpRequest.Email
	newUser.Password = string(as.al.CreateHashPassword(signUpRequest.Password))

	if err := as.ur.CreateUser(&newUser); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "ユーザー登録処理に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.User{}, err
	}

	return newUser, nil
}

// レスポンス送信

func (as *authService) SendAuthResponse(w http.ResponseWriter, user *models.User, code int) {
	// jwtトークン作成
	token, err := as.jl.CreateJwtToken(user)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "トークン作成に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(errMessage), statusCode)
	}

	var res models.AuthResponse
	res.Token = token
	res.User.BaseModel.ID = user.ID
	res.User.BaseModel.CreatedAt = user.CreatedAt
	res.User.BaseModel.UpdatedAt = user.UpdatedAt
	res.User.BaseModel.DeletedAt = user.DeletedAt
	res.User.Name = user.Name
	res.User.Email = user.Email

	body, _ := json.Marshal(res)

	as.rl.SendResponse(w, body, code)
}
