package dashboard

import (
	"context"
	"try-graphql/internal/app/entity"
)

type Inport interface {
	Execute(ctx context.Context, planetUUID string) (*entity.Planet, []entity.PlanetDistrict, []entity.PlanetBuilding, []entity.PlanetaryFeature, error)
}
