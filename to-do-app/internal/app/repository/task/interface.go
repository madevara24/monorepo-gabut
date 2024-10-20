package task

import (
	"context"
	"to-do-app/internal/app/entities"
	"to-do-app/internal/pkg/datasource"
)

type IRepository interface {
	Create(context.Context, entities.Task) error
	// Update(ctx context.Context, task entities.Task) (entities.Task, error)
	// Delete(ctx context.Context, uuid string) error
	GetByUUID(ctx context.Context, uuid string) (entities.Task, error)
	// GetAll(ctx context.Context) ([]entities.Task, error)
}

type repo struct {
	datasource *datasource.DataSource
}

func NewRepo(datasource *datasource.DataSource) IRepository {
	return &repo{
		datasource: datasource,
	}
}
