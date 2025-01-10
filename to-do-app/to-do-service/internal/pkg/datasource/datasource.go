package datasource

import (
	"fmt"
	"log"
	"time"
	"to-do-service/config"

	"github.com/madevara24/go-common/database"

	"github.com/jmoiron/sqlx"
)

type DataSource struct {
	Postgre *sqlx.DB
}

func NewDataSource() *DataSource {
	postgresClient := database.NewConfiguration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Get().DBUsername,
		config.Get().DBPassword,
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBName,
	), "cultivation-sqlx")

	postgresDB, err := sqlx.Connect("postgres", postgresClient.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	postgresDB.SetMaxIdleConns(config.Get().DBMaxIdleConn)
	postgresDB.SetMaxOpenConns(config.Get().DBMaxConn)
	postgresDB.SetConnMaxLifetime(time.Duration(config.Get().DBMaxTTLConn) * time.Second)

	return &DataSource{
		Postgre: postgresDB,
	}
}
