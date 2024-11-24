package app

import (
	"to-do-app/internal/app/repository/task"
	"to-do-app/internal/app/usecase/healthcheck"
	"to-do-app/internal/app/usecase/task/create"
	"to-do-app/internal/app/usecase/task/getall"
	"to-do-app/internal/app/usecase/task/getbyuuid"
	"to-do-app/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport

	// TASK
	CreateTaskInport    create.Inport
	GetTaskByUUIDInport getbyuuid.Inport
	GetAllTaskInport    getall.Inport
}

func NewContainer(datasource *datasource.DataSource) *Container {

	taskRepo := task.NewRepo(datasource)
	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),

		// TASK
		CreateTaskInport:    create.NewUsecase(taskRepo),
		GetTaskByUUIDInport: getbyuuid.NewUsecase(taskRepo),
		GetAllTaskInport:    getall.NewUsecase(taskRepo),
	}
}
