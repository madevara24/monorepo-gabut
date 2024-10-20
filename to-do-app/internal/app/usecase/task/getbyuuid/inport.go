package getbyuuid

import (
	"context"
	"to-do-app/internal/app/entities"
)

type Inport interface {
	Execute(ctx context.Context, req InportRequest) (InportResponse, error)
}

type InportRequest struct {
	UUID string
}

type InportResponse struct {
	entities.Task
}
