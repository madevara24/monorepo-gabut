package create

import (
	"context"
	"to-do-service/internal/app/entities"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type Inport interface {
	Execute(ctx context.Context, payload InportRequest) error
}

type InportRequest struct {
	Title       string              `json:"title" validate:"required"`
	Description null.String         `json:"description"`
	Deadline    null.Time           `json:"deadline"`
	Status      entities.TaskStatus `json:"status" validate:"required"`
	CreatedBy   string              `json:"created_by" validate:"required"`
}

func (i *InportRequest) MapIntoTask() (entities.Task, error) {
	var task entities.Task = entities.Task{
		UUID:        uuid.NewString(),
		Title:       i.Title,
		Description: i.Description,
		Deadline:    i.Deadline,
		Status:      i.Status,
		CreatedBy:   i.CreatedBy,
		UpdatedBy:   i.CreatedBy,
	}
	if err := task.Validate(); err != nil {
		return entities.Task{}, err
	}
	return task, nil
}
