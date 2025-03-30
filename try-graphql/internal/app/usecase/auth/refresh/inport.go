package refresh

import (
	"context"
	"try-graphql/internal/app/entity"
)

type Inport interface {
	Execute(ctx context.Context, req InportRequest) (entity.Token, error)
}

type InportRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
