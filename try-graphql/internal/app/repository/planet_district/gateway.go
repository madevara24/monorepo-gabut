package planet_district

import (
	"context"
	"try-graphql/internal/app/entity"
)

type IRepo interface {
	GetByPlanetUUID(ctx context.Context, planetUUID string) ([]entity.PlanetDistrict, error)
}
