package rest

import (
	"context"
	"try-graphql/internal/app"
	"try-graphql/internal/app/delivery/rest/auth"
	"try-graphql/internal/app/delivery/rest/healthcheck"
	"try-graphql/internal/app/delivery/rest/middleware"
	"try-graphql/internal/app/delivery/rest/user"
	"try-graphql/internal/pkg/datasource"

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
	h.router.Use(gin.Recovery())

	v1 := h.router.Group("/v1")

	// PING
	v1.GET("/health", healthcheck.HealthCheckHandler(h.container.HealthCheckInport))

	// AUTH
	authRouter := v1.Group("/auth")
	authRouter.POST("/login", auth.Login(h.container.AuthLoginInport))
	authRouter.POST("/refresh", auth.Refresh(h.container.AuthRefreshInport))

	// USER
	userRouter := v1.Group("/user")
	userRouter.POST("/register", user.Register(h.container.UserRegisterInport))

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Add protected routes here
	}
}
