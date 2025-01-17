package rest

import (
	"context"
	"to-do-service/internal/app"
	"to-do-service/internal/app/delivery/rest/healthcheck"
	"to-do-service/internal/app/delivery/rest/task"
	"to-do-service/internal/pkg/datasource"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	datasource *datasource.DataSource
	container  *app.Container
}

func NewRouter(ctx context.Context, router *gin.Engine, datasource *datasource.DataSource, container *app.Container) *Router {
	return &Router{
		router:     router,
		datasource: datasource,
		container:  container,
	}
}

func (h *Router) RegisterRouter() {

	h.router.Use(SetTDRMiddleware())
	h.router.Use(gin.Recovery())
	// PING
	h.router.GET("/health", healthcheck.HealthCheckHandler(h.container.HealthCheckInport))

	h.router.POST("/tasks", task.Create(h.container.CreateTaskInport))
	h.router.GET("/tasks", task.GetAll(h.container.GetAllTaskInport))
	h.router.GET("/tasks/:uuid", task.GetByUUID(h.container.GetTaskByUUIDInport))
}
