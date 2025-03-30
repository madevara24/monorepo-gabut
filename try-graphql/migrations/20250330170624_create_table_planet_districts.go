package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250330170624",
		Up:      mig_20250330170624_create_table_planet_districts_up,
		Down:    mig_20250330170624_create_table_planet_districts_down,
	})
}

func mig_20250330170624_create_table_planet_districts_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS planet_districts (
			uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			planet_uuid UUID NOT NULL REFERENCES planets(uuid),
			type VARCHAR(20) NOT NULL,
			level INTEGER NOT NULL
		);
	`)
	return err
}

func mig_20250330170624_create_table_planet_districts_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS planet_districts;
	`)
	return err
}
