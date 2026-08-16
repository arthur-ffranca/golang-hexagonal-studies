package usecase

import (
	"context"

	"fastapi2/hexagonal/application/ports"
	"fastapi2/hexagonal/internal/domain"
)

type CreateUserInput struct {
	ID    string
	Name  string
	Email string
}

type CreateUserUseCase struct {
	repo ports.UserRepository
}

func NewCreateUserUseCase(repo ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (domain.User, error) {
	user, err := domain.NewUser(input.ID, input.Name, input.Email)
	if err != nil {
		return domain.User{}, err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uc *CreateUserUseCase) FindByID(ctx context.Context, id string) (domain.User, error) {
	return uc.repo.FindByID(ctx, id)
}