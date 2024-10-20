package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20241017170716",
		Up:      mig_20241017170716_create_task_table_up,
		Down:    mig_20241017170716_create_task_table_down,
	})
}

func mig_20241017170716_create_task_table_up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		-- SAFELY CREATE TYPE
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tasks_status') THEN
				CREATE TYPE tasks_status AS ENUM ('PENDING', 'IN_PROGRESS', 'COMPLETED');
			END IF;
		END $$;

		CREATE TABLE IF NOT EXISTS tasks (
			uuid uuid PRIMARY KEY,
			deadline timestamptz NULL,
			description text NULL,
			status tasks_status NOT NULL DEFAULT 'PENDING',
			created_by varchar(100) NOT NULL,
			updated_by varchar(100) NOT NULL,
			deleted_by varchar(100) NULL,
			created_at timestamptz NOT NULL DEFAULT now(),
			updated_at timestamptz NOT NULL DEFAULT now(),
			deleted_at timestamptz NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func mig_20241017170716_create_task_table_down(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS tasks;
			
		DO $$ BEGIN
			IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tasks_status') THEN
				DROP TYPE tasks_status;
			END IF;
		END $$;
	`)
	if err != nil {
		return err
	}
	return nil
}
