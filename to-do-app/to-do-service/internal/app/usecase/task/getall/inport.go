package getall

import (
	"context"
	"to-do-service/internal/app/entities"
)

type Inport interface {
	Execute(ctx context.Context) (InportResponse, error)
}

type InportResponse struct {
	Tasks []entities.Task `json:"tasks"`
}
