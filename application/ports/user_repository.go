package ports

import (
	"context"

	"fastapi2/hexagonal/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByID(ctx context.Context, id string) (domain.User, error)
}
