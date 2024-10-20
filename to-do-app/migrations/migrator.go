package migrations

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/madevara24/go-common/array"

	"github.com/jmoiron/sqlx"
)

var migrator = &Migrator{
	Versions:   []string{},
	Migrations: map[string]*Migration{},
}

// Init : initialize schema migrations
func Init(db *sqlx.DB) (*Migrator, error) {
	migrator.db = db

	// Create schema_migrations table to remember which migrations were executed.
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS schema_migrations (version varchar(255));"); err != nil {
		fmt.Println("Unable to create schema_migrations table", err)
		return migrator, err
	}

	rows, err := db.Query("SELECT version FROM schema_migrations;")
	if err != nil {
		return migrator, err
	}

	defer rows.Close()

	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return migrator, err
		}

		if migrator.Migrations[version] != nil {
			migrator.Migrations[version].done = true
		}
	}

	return migrator, nil
}

func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./migrations/template.txt"))
	if err := t.Execute(&out, in); err != nil {
		return errors.New("Unable to execute template: " + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("./migrations/%s_%s.go", version, name))

	if err != nil {
		return errors.New("Unable to create migration file: " + err.Error())
	}

	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file: " + err.Error())
	}

	fmt.Println("Generated new migration file...", f.Name())
	return nil
}

// AddMigration : add new migration version
func (m *Migrator) AddMigration(mg *Migration) {
	// Add the migration to the hash with version as key
	m.Migrations[mg.Version] = mg

	// Insert version into versions array using insertion sort
	index := 0
	for index < len(m.Versions) {
		if m.Versions[index] > mg.Version {
			break
		}

		index++
	}

	m.Versions = append(m.Versions, mg.Version)
	copy(m.Versions[index+1:], m.Versions[index:])
	m.Versions[index] = mg.Version
}

// Up .
func (m *Migrator) Up(step int) error {
	tx, err := m.db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	count := 0

	for _, v := range m.Versions {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if mg.done {
			continue
		}

		fmt.Println("Running migration", mg.Version)
		if err := mg.Up(tx); err != nil {
			tx.Rollback()
			return err
		}

		if _, err := tx.Exec("INSERT INTO schema_migrations VALUES($1)", mg.Version); err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println("Finished running migration", mg.Version)

		count++
	}

	tx.Commit()

	return nil
}

// Down .
func (m *Migrator) Down(step int) error {
	tx, err := m.db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	count := 0
	for _, v := range array.Reverse(m.Versions) {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if !mg.done {
			continue
		}

		fmt.Println("Reverting migration", mg.Version)
		if err := mg.Down(tx); err != nil {
			tx.Rollback()
			return err
		}

		if _, err := tx.Exec("DELETE FROM schema_migrations WHERE version = $1", mg.Version); err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println("Finished reverting migration", mg.Version)

		count++
	}

	tx.Commit()

	return nil
}

// MigrationStatus .
func (m *Migrator) MigrationStatus() error {
	for _, v := range m.Versions {
		mg := m.Migrations[v]

		if mg.done {
			fmt.Printf("%s", fmt.Sprintf("Migration %s... completed\n", v))
		} else {
			fmt.Printf("%s", fmt.Sprintf("Migration %s... pending\n", v))
		}
	}

	return nil
}
