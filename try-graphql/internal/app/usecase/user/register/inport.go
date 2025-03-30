package register

import (
	"context"
	"time"
	"try-graphql/internal/app/entity"
	codebaseErrors "try-graphql/internal/pkg/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Inport interface {
	Execute(ctx context.Context, req InportRequest) error
}

type InportRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (i *InportRequest) MapIntoUser() (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, codebaseErrors.ErrPasswordHashFailed
	}

	return entity.User{
		UUID:      uuid.New().String(),
		Email:     i.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: i.Email,
		UpdatedBy: i.Email,
	}, nil
}
