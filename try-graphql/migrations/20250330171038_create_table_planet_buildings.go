package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250330171038",
		Up:      mig_20250330171038_create_table_planet_buildings_up,
		Down:    mig_20250330171038_create_table_planet_buildings_down,
	})
}

func mig_20250330171038_create_table_planet_buildings_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS planet_buildings (
			uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			district_uuid UUID NOT NULL REFERENCES planet_districts(uuid),
			type VARCHAR(20) NOT NULL,
			level INTEGER NOT NULL
		);
	`)
	return err
}

func mig_20250330171038_create_table_planet_buildings_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS planet_buildings;
	`)
	return err
}
