package cmd

import (
	"fmt"
	"log"
	"os"
	"to-do-app/migrations"

	"github.com/madevara24/go-common/database"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

func initDB() *sqlx.DB {
	configPostgre := database.NewConfiguration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	), os.Getenv("DB_SQLX_KEY"))
	err := database.NewPostgres(configPostgre)
	if err != nil {
		panic(err)
	}

	return database.GetSqlxClient(os.Getenv("DB_SQLX_KEY"))
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migration file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migrations.Create(name); err != nil {
			log.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			log.Println("Unable to read flag `step`")
			return
		}

		// setup database connection
		db := initDB()

		migrator, err := migrations.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		log.Println("Running migration...")
		err = migrator.Up(step)
		if err != nil {
			log.Println("Unable to run `up` migrations", err.Error())
			return
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			log.Println("Unable to read flag `step`")
			return
		}

		// setup database connection
		db := initDB()

		migrator, err := migrations.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		err = migrator.Down(step)
		if err != nil {
			log.Println("Unable to run `down` migrations")
			return
		}
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()

		migrator, err := migrations.Init(db)
		if err != nil {
			log.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			log.Println("Unable to fetch migration status")
			return
		}
	},
}

func init() {
	// Add "--name" flag to "create" command
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")
	// set flag "name" required
	_ = migrateCreateCmd.MarkFlagRequired("name")

	// Add "--step" flag to both "up" and "down" command
	migrateUpCmd.Flags().IntP("step", "s", 0, "Number of migration to execute")
	migrateDownCmd.Flags().IntP("step", "s", 0, "Number of migration to execute")
}
