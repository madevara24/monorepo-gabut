package migrations

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Migration information & command struct
type Migration struct {
	Version string
	Up      func(*sql.Tx) error
	Down    func(*sql.Tx) error

	done bool
}

// Migrator : collection of migration
type Migrator struct {
	db         *sqlx.DB
	Versions   []string
	Migrations map[string]*Migration
}
