package graphql

import (
	"context"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/app/usecase/planet/dashboard"
)

type Resolver struct {
	planetDashboard dashboard.Inport
}

func New(planetDashboard dashboard.Inport) *Resolver {
	return &Resolver{
		planetDashboard: planetDashboard,
	}
}

func (r *Resolver) PlanetDashboard(ctx context.Context, planetUUID string) (*Planet, error) {
	planet, districts, buildings, features, err := r.planetDashboard.Execute(ctx, planetUUID)
	if err != nil {
		return nil, err
	}

	// Map the data to GraphQL types
	return &Planet{
		UUID:      planet.UUID,
		Name:      planet.Name,
		Size:      planet.Size,
		Type:      planet.Type,
		Districts: mapDistricts(districts, buildings),
		Features:  mapFeatures(features),
	}, nil
}

func mapDistricts(districts []entity.PlanetDistrict, buildings []entity.PlanetBuilding) []*PlanetDistrict {
	districtMap := make(map[string]*PlanetDistrict)

	// First, create district entries
	for _, d := range districts {
		districtMap[d.UUID] = &PlanetDistrict{
			UUID:  d.UUID,
			Type:  d.Type,
			Level: d.Level,
		}
	}

	// Then, add buildings to their respective districts
	for _, b := range buildings {
		if district, exists := districtMap[b.DistrictUUID]; exists {
			district.Buildings = append(district.Buildings, &PlanetBuilding{
				UUID:  b.UUID,
				Type:  b.Type,
				Level: b.Level,
			})
		}
	}

	// Convert map to slice
	result := make([]*PlanetDistrict, 0, len(districtMap))
	for _, d := range districtMap {
		result = append(result, d)
	}
	return result
}

func mapFeatures(features []entity.PlanetaryFeature) []*PlanetaryFeature {
	result := make([]*PlanetaryFeature, len(features))
	for i, f := range features {
		result[i] = &PlanetaryFeature{
			UUID: f.UUID,
			Type: f.Type,
		}
	}
	return result
}
