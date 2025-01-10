package create

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

func (i interactor) Execute(ctx context.Context, req InportRequest) error {
	task, err := req.MapIntoTask()
	if err != nil {
		return err
	}
	err = i.taskRepo.Create(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
