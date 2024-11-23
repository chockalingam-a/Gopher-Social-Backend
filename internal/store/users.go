package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	/* IsActive  bool     `json:"is_active"`
	RoleID    int64    `json:"role_id"` */
	/* Role      Role     `json:"role"` */
}

type password struct {
	text *string
	hash []byte
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO (username, email, password)
		VALUES($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Username,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
