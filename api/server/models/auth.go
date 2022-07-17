package models

/*
ログインパラメータ
*/
type SignInRequest struct {
	Email    string `gorm:"varchar(255);not null;unique" json:"email,omitempty" example:"test@example.com"`
	Password string `gorm:"varchar(255);not null" json:"password,omitempty" example:"password"`
}

/*
会員登録パラメータ
*/
type SignUpRequest struct {
	Name     string `gorm:"varchar(255);not null" json:"name,omitempty" example:"test user"`
	Email    string `gorm:"varchar(255);not null;unique" json:"email,omitempty" example:"test@example.com"`
	Password string `gorm:"varchar(255);not null" json:"password,omitempty" example:"password"`
}

type AuthResponse struct {
	Token string           `json:"token" example:"ajiji1j98a"`
	User  BaseUserResponse `json:"user"`
}
