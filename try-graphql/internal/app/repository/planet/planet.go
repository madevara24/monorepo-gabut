package planet

import (
	"context"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/pkg/datasource"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *datasource.DataSource) IRepo {
	return &repo{
		db: db.Postgres,
	}
}

func (r *repo) GetByUUID(ctx context.Context, planetUUID string) (*entity.Planet, error) {
	var planet entity.Planet
	err := r.db.GetContext(ctx, &planet, "SELECT * FROM planets WHERE uuid = $1", planetUUID)
	if err != nil {
		return nil, err
	}
	return &planet, nil
}
