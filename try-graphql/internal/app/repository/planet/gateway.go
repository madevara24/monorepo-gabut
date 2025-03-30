package planet

import (
	"context"
	"try-graphql/internal/app/entity"
)

type IRepo interface {
	GetByUUID(ctx context.Context, planetUUID string) (*entity.Planet, error)
}
