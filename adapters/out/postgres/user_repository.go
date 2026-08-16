package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"fastapi2/hexagonal/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	query := `
		INSERT INTO users (id, name, email)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query, user.ID, user.Name, user.Email)
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (domain.User, error) {
	query := `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	var user domain.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
