package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250330170137",
		Up:      mig_20250330170137_create_table_planetary_features_up,
		Down:    mig_20250330170137_create_table_planetary_features_down,
	})
}

func mig_20250330170137_create_table_planetary_features_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS planetary_features (
			uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			planet_uuid UUID NOT NULL REFERENCES planets(uuid),
			type VARCHAR(10) NOT NULL
		);
	`)
	return err
}

func mig_20250330170137_create_table_planetary_features_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS planetary_features;
	`)
	return err
}
