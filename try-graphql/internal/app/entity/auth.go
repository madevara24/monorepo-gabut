package entity

import (
	"net/http"

	goCommonError "github.com/madevara24/go-common/errors"
)

const (
	// error code
	ERR_CODE_INVALID_CREDENTIALS = "AUTH_001"
	ERR_CODE_TOKEN_EXPIRED       = "AUTH_002"
	ERR_CODE_TOKEN_INVALID       = "AUTH_003"
)

var (
	ErrInvalidCredentials = goCommonError.NewErr(http.StatusUnauthorized, ERR_CODE_INVALID_CREDENTIALS, "invalid credentials")
	ErrTokenExpired       = goCommonError.NewErr(http.StatusUnauthorized, ERR_CODE_TOKEN_EXPIRED, "token has expired")
	ErrTokenInvalid       = goCommonError.NewErr(http.StatusUnauthorized, ERR_CODE_TOKEN_INVALID, "invalid token")
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
