package logic

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthLogic interface {
	GetUserIdFromToken(r *http.Request) (int, error)
	CreateHashPassword(password string) []byte
}

type authLogic struct {
}

func NewAuthLogic() AuthLogic {
	return &authLogic{}
}

func (al *authLogic) GetUserIdFromToken(r *http.Request) (int, error) {
	clientToken := r.Header.Get("Authorization")
	if clientToken == "" {
		return 0, errors.New("not token")
	}

	extractedToken := strings.Split(clientToken, "Bearer ")
	secretKey := os.Getenv("JWT_SECRET")

	var parseToken string
	// Bearerがなかった場合の対処 (OpenAPI Ver.2)
	if len(extractedToken) == 1 {
		parseToken = extractedToken[0]
	} else {
		parseToken = extractedToken[1]
		if parseToken == "" {
			return 0, errors.Errorf("トークンが空文字です。")
		}
	}

	token, err := jwt.Parse(
		parseToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Errorf("トークンをjwtにparseできません。")
			}
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, claimOk := token.Claims.(jwt.MapClaims)
	if !claimOk || !token.Valid {
		return 0, errors.New("id type not match")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("id type not match")
	}

	return int(userId), nil
}

func (al *authLogic) CreateHashPassword(password string) []byte {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashPassword
}
