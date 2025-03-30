package login

import (
	"context"
	"try-graphql/internal/app/entity"
)

type Inport interface {
	Execute(ctx context.Context, req InportRequest) (entity.Token, error)
}

type InportRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
