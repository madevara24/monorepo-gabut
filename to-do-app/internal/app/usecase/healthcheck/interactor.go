package healthcheck

import (
	"context"
	"time"
	"to-do-app/config"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	postgres *sqlx.DB
}

// NewUsecase --
func NewUsecase(postgres *sqlx.DB) Inport {
	return interactor{
		postgres: postgres,
	}
}

func (i interactor) Execute(ctx context.Context) InportResponse {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	postgres := true
	if err := i.postgres.PingContext(ctx); err != nil {
		postgres = false
	}

	message := "OK"
	if postgres != true {
		message = "NOT OK"
	}

	return InportResponse{
		Name:    "Service Cultivation",
		Message: message,
		Status: Status{
			Postgres: postgres,
		},
		Version: config.Get().AppVersion,
	}
}
