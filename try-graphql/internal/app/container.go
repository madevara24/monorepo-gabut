package app

import (
	"try-graphql/internal/app/repository/user"
	"try-graphql/internal/app/usecase/auth/login"
	"try-graphql/internal/app/usecase/auth/refresh"
	"try-graphql/internal/app/usecase/healthcheck"
	"try-graphql/internal/app/usecase/user/register"
	"try-graphql/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport

	// USER
	UserRegisterInport register.Inport

	// AUTH
	AuthLoginInport   login.Inport
	AuthRefreshInport refresh.Inport
}

func NewContainer(datasource *datasource.DataSource) *Container {
	userRepo := user.NewRepo(datasource)
	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),

		// USER
		UserRegisterInport: register.NewUsecase(userRepo),

		// AUTH
		AuthLoginInport:   login.NewUsecase(userRepo),
		AuthRefreshInport: refresh.NewUsecase(),
	}
}
