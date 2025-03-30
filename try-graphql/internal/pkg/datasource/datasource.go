package datasource

import (
	"fmt"
	"log"
	"time"
	"try-graphql/config"

	"github.com/madevara24/go-common/database"
	"github.com/madevara24/go-common/mapper"

	"github.com/jmoiron/sqlx"
)

type DataSource struct {
	Postgres *sqlx.DB
	Mapper   mapper.IMapper
}

func NewDataSource() *DataSource {
	postgresClient := database.NewConfiguration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Get().DBUsername,
		config.Get().DBPassword,
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBName,
	), config.Get().DBSqlxKey)

	postgresDB, err := sqlx.Connect("postgres", postgresClient.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	postgresDB.SetMaxIdleConns(config.Get().DBMaxIdleConn)
	postgresDB.SetMaxOpenConns(config.Get().DBMaxConn)
	postgresDB.SetConnMaxLifetime(time.Duration(config.Get().DBMaxTTLConn) * time.Second)

	return &DataSource{
		Postgres: postgresDB,
		Mapper:   mapper.NewPostgresMapper(),
	}
}
