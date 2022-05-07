package logic

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IAuthLogic interface {
	GetUserIdFromToken(r *http.Request) (int, error)
	CreateHashPassword(password string) []byte
}

type AuthLogic struct {
}

func NewAuthLogic() *AuthLogic {
	return &AuthLogic{}
}

func (al *AuthLogic) GetUserIdFromToken(r *http.Request) (int, error) {
	clientToken := r.Header.Get("Authorization")
	if clientToken == "" {
		return 0, errors.New("not token")
	}

	extractToken := strings.Split(clientToken, "Bearer ")
	secretKey := os.Getenv("JWT_KEY")

	if extractToken[1] == "" {
		return 0, errors.Errorf("トークンが空文字です。")
	}

	token, err := jwt.Parse(extractToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("トークンをjwtにparseできません。")
		}
		return []byte(secretKey), nil
	})
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

func (al *AuthLogic) CreateHashPassword(password string) []byte {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashPassword
}
