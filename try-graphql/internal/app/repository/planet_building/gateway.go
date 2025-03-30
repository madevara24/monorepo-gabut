package planet_building

import (
	"context"
	"try-graphql/internal/app/entity"
)

type IRepo interface {
	GetByPlanetUUID(ctx context.Context, planetUUID string) ([]entity.PlanetBuilding, error)
}
