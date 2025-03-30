package planet_feature

import (
	"context"
	"try-graphql/internal/app/entity"
)

type IRepo interface {
	GetByPlanetUUID(ctx context.Context, planetUUID string) ([]entity.PlanetaryFeature, error)
}
