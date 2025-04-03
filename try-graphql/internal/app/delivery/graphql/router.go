package graphql

import (
	"context"
	"try-graphql/graph/generated"
	"try-graphql/internal/app"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router    *gin.Engine
	container *app.Container
}

func NewRouter(ctx context.Context, router *gin.Engine, container *app.Container) *Router {
	return &Router{
		router:    router,
		container: container,
	}
}

func (h *Router) RegisterRouter() {
	// Create a new resolver with the needed dependencies
	resolver := New(h.container.PlanetDashboardInport)

	// Create a GraphQL handler
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
	)

	// Create a GraphQL playground handler
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")

	v1 := h.router.Group("/graphql/v1")

	// Register routes
	v1.GET("/playground", gin.WrapH(playgroundHandler))
	v1.POST("/query", gin.WrapH(srv))
}
