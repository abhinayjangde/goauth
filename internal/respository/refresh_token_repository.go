package respository

import (
	"context"

	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Save(token *model.RefreshToken) error {

	query := `
	INSERT INTO refresh_tokens
	(user_id, token, expires_at)
	VALUES ($1,$2,$3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		token.UserId,
		token.Token,
		token.ExpiresAt,
	)

	return err
}

func (r *RefreshTokenRepository) Find(token string) (*model.RefreshToken, error) {

	query := `
	SELECT
		id,
		user_id,
		token,
		expires_at,
		created_at
	FROM refresh_tokens
	WHERE token=$1
	`

	rt := &model.RefreshToken{}

	err := r.db.QueryRow(
		context.Background(),
		query,
		token,
	).Scan(
		&rt.ID,
		&rt.UserId,
		&rt.Token,
		&rt.ExpiresAt,
		&rt.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (r *RefreshTokenRepository) Delete(token string) error {

	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM refresh_tokens
		 WHERE token=$1`,
		token,
	)

	return err
}
