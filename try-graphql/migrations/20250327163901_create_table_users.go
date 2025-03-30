package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20250327163901",
		Up:      mig_20250327163901_create_table_users_up,
		Down:    mig_20250327163901_create_table_users_down,
	})
}

func mig_20250327163901_create_table_users_up(tx *sql.Tx) error {

	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP NULL,
			created_by VARCHAR(255) NULL,
			updated_by VARCHAR(255) NULL,
			deleted_by VARCHAR(255) NULL
		);
	`)
	return err
}

func mig_20250327163901_create_table_users_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS users;
	`)
	return err
}
