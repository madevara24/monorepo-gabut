package dashboard

import (
	"context"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/app/repository/planet"
	planetBuilding "try-graphql/internal/app/repository/planet_building"
	planetDistrict "try-graphql/internal/app/repository/planet_district"
	planetFeature "try-graphql/internal/app/repository/planet_feature"
)

type interactor struct {
	planetRepo   planet.IRepo
	districtRepo planetDistrict.IRepo
	buildingRepo planetBuilding.IRepo
	featureRepo  planetFeature.IRepo
}

func New(
	planetRepo planet.IRepo,
	districtRepo planetDistrict.IRepo,
	buildingRepo planetBuilding.IRepo,
	featureRepo planetFeature.IRepo,
) Inport {
	return &interactor{
		planetRepo:   planetRepo,
		districtRepo: districtRepo,
		buildingRepo: buildingRepo,
		featureRepo:  featureRepo,
	}
}

func (i *interactor) Execute(ctx context.Context, planetUUID string) (*entity.Planet, []entity.PlanetDistrict, []entity.PlanetBuilding, []entity.PlanetaryFeature, error) {
	planet, err := i.planetRepo.GetByUUID(ctx, planetUUID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	districts, err := i.districtRepo.GetByPlanetUUID(ctx, planetUUID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	buildings, err := i.buildingRepo.GetByPlanetUUID(ctx, planetUUID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	features, err := i.featureRepo.GetByPlanetUUID(ctx, planetUUID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return planet, districts, buildings, features, nil
}
