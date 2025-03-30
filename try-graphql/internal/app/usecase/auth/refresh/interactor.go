package refresh

import (
	"context"
	"time"
	"try-graphql/config"
	"try-graphql/internal/app/entity"

	"github.com/golang-jwt/jwt/v5"
)

type interactor struct{}

func NewUsecase() Inport {
	return &interactor{}
}

func (i *interactor) Execute(ctx context.Context, req InportRequest) (entity.Token, error) {
	// Parse and validate the refresh token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().JWTSecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return entity.Token{}, entity.ErrTokenExpired
		}
		return entity.Token{}, entity.ErrTokenInvalid
	}

	// Extract user ID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return entity.Token{}, entity.ErrTokenInvalid
	}

	userID := claims["sub"].(string)

	// Generate new access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(config.Get().JWTExpiryHours)).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(config.Get().JWTSecretKey))
	if err != nil {
		return entity.Token{}, err
	}

	// Generate new refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(config.Get().JWTRefreshHours)).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(config.Get().JWTSecretKey))
	if err != nil {
		return entity.Token{}, err
	}

	return entity.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(config.Get().JWTExpiryHours * 3600),
		TokenType:    "Bearer",
	}, nil
}
