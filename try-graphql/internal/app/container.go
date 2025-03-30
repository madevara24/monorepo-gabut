package app

import (
	"try-graphql/internal/app/delivery/graphql"
	"try-graphql/internal/app/repository/planet"
	planetBuilding "try-graphql/internal/app/repository/planet_building"
	planetDistrict "try-graphql/internal/app/repository/planet_district"
	planetFeature "try-graphql/internal/app/repository/planet_feature"
	"try-graphql/internal/app/usecase/healthcheck"
	"try-graphql/internal/app/usecase/planet/dashboard"
	"try-graphql/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport

	// GraphQL
	GraphQLResolver *graphql.Resolver
}

func NewContainer(datasource *datasource.DataSource) *Container {
	// Repositories
	planetRepo := planet.New(datasource)
	districtRepo := planetDistrict.New(datasource)
	buildingRepo := planetBuilding.New(datasource)
	featureRepo := planetFeature.New(datasource)

	// Usecases
	planetDashboard := dashboard.New(planetRepo, districtRepo, buildingRepo, featureRepo)

	// GraphQL
	graphqlResolver := graphql.New(planetDashboard)

	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgres),

		// GraphQL
		GraphQLResolver: graphqlResolver,
	}
}
