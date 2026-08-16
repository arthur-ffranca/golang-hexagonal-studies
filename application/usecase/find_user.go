package usecase

import (
	"context"

	"fastapi2/hexagonal/application/ports"
	"fastapi2/hexagonal/internal/domain"
)

type FindUserInput struct {
	ID string
}

type FindUserUseCase struct {
	repo ports.UserRepository
}

func NewFindUserUseCase(repo ports.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		repo: repo,
	}
}

func (uc *FindUserUseCase) Execute(ctx context.Context, input FindUserInput) (domain.User, error) {
	if input.ID == "" {
		return domain.User{}, domain.ErrInvalidUser
	}

	return uc.repo.FindByID(ctx, input.ID)
}