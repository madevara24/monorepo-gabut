package getall

import (
	"context"
	"to-do-app/internal/app/repository/task"
)

type interactor struct {
	taskRepo task.IRepository
}

func NewUsecase(taskRepo task.IRepository) Inport {
	return interactor{
		taskRepo: taskRepo,
	}
}

func (i interactor) Execute(ctx context.Context) (InportResponse, error) {
	tasks, err := i.taskRepo.GetAll(ctx)
	if err != nil {
		return InportResponse{}, err
	}
	return InportResponse{
		Tasks: tasks,
	}, nil
}
