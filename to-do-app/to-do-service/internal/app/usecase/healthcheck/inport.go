package healthcheck

import (
	"context"
)

type Inport interface {
	Execute(context.Context) InportResponse
}

type InportResponse struct {
	Message string `json:"message"`
	Status  Status `json:"status"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Status struct {
	Postgres bool `json:"postgres"`
}
