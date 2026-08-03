package respository

import (
	"context"
	"fmt"

	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRespository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) error {

	// check if user already exists
	_, userNotFoundErr := r.FindByEmail(user.Email)
	if userNotFoundErr == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	query := `
		INSERT INTO users(name, email, password)
		VALUES($1,$2,$3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
	)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email=$1
	`

	user := &model.User{}

	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
