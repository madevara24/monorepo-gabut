package planet_building

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

func (r *repo) GetByPlanetUUID(ctx context.Context, planetUUID string) ([]entity.PlanetBuilding, error) {
	var buildings []entity.PlanetBuilding
	err := r.db.SelectContext(ctx, &buildings, "SELECT * FROM planet_buildings WHERE district_uuid IN (SELECT uuid FROM planet_districts WHERE planet_uuid = $1)", planetUUID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return buildings, nil
}
