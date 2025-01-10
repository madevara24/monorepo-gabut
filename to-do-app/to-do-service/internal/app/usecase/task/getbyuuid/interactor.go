package getbyuuid

import (
	"context"
	"to-do-service/internal/app/repository/task"
)

type interactor struct {
	taskRepo task.IRepository
}

func NewUsecase(taskRepo task.IRepository) Inport {
	return interactor{
		taskRepo: taskRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	task, err := i.taskRepo.GetByUUID(ctx, req.UUID)
	if err != nil {
		return InportResponse{}, err
	}
	return InportResponse{
		Task: task,
	}, nil
}
