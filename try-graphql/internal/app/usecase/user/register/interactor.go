package register

import (
	"context"
	"try-graphql/internal/app/entity"
	"try-graphql/internal/app/repository/user"
)

type interactor struct {
	userRepo user.IRepo
}

func NewUsecase(userRepo user.IRepo) Inport {
	return interactor{
		userRepo: userRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) error {

	user, err := req.MapIntoUser()
	if err != nil {
		return err
	}

	err = user.Validate()
	if err != nil {
		return err
	}

	exissting, err := i.userRepo.FindByEmail(ctx, user.Email)
	if err != nil && err != entity.ErrUserNotFound {
		return err
	}

	if exissting.Email != "" {
		return entity.ErrEmailExists
	}

	err = i.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
