package login

import (
	"context"
	"time"
	"try-graphql/config"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/app/repository/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type interactor struct {
	userRepo user.IRepo
}

func NewUsecase(userRepo user.IRepo) Inport {
	return &interactor{
		userRepo: userRepo,
	}
}

func (i *interactor) Execute(ctx context.Context, req InportRequest) (entity.Token, error) {
	user, err := i.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return entity.Token{}, entity.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return entity.Token{}, entity.ErrInvalidCredentials
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UUID,
		"exp": time.Now().Add(time.Hour * time.Duration(config.Get().JWTExpiryHours)).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(config.Get().JWTSecretKey))
	if err != nil {
		return entity.Token{}, err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UUID,
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
