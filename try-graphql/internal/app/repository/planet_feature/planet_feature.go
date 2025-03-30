package planet_feature

import (
	"context"
	"database/sql"
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

func (r *repo) GetByPlanetUUID(ctx context.Context, planetUUID string) ([]entity.PlanetaryFeature, error) {
	var features []entity.PlanetaryFeature
	err := r.db.SelectContext(ctx, &features, "SELECT * FROM planetary_features WHERE planet_uuid = $1", planetUUID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return features, nil
}
