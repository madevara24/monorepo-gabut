package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250330165130",
		Up:      mig_20250330165130_create_table_planets_up,
		Down:    mig_20250330165130_create_table_planets_down,
	})
}

func mig_20250330165130_create_table_planets_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DO $$
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'planet_type') THEN
					CREATE TYPE planet_type AS ENUM (
						'TROPICAL',
						'OCEAN',
						'CONTINENTAL',
						'SAVANNA',
						'DESERT',
						'ARID',
						'TUNDRA',
						'ARCTIC',
						'ALPINE'
					);
				END IF;

				CREATE TABLE IF NOT EXISTS planets (
					uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					name VARCHAR(255) NOT NULL,
					size INTEGER NOT NULL,
					type planet_type NOT NULL
				);
			END
		$$;
	`)
	return err
}

func mig_20250330165130_create_table_planets_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS planets;

		DO $$
			BEGIN
				IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'planet_type') THEN
					DROP TYPE planet_type;
				END IF;
			END
		$$;
	`)
	return err
}
