package user

import (
	"context"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/pkg/datasource"
)

type IRepo interface {
	Create(ctx context.Context, req entity.User) error
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}

type repo struct {
	datasource *datasource.DataSource
}

func NewRepo(datasource *datasource.DataSource) IRepo {
	return &repo{
		datasource: datasource,
	}
}
