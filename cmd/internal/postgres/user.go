package postgres

import (
	"context"
	"new-go-project/cmd/service"
)

func NewUserStore(
	db *Client,
) service.UserStore {
	return &userStore{
		db: db,
	}
}

type userStore struct {
	db *Client
}

// transactional ke database:
// BEGIN
// EXECUTION
// COMMIT

func (s *userStore) CreateUser(ctx context.Context, user *service.User) error {
	tx, err := s.db.Leader.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (
		id,
		first_name,
		middle_name,
		last_name,
		type,
		status
	) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.ExecContext(ctx, query,
		user.ID,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.Type,
		user.Status,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *userStore) GetUsers(ctx context.Context) ([]service.User, error) {
	query := `SELECT
		id,
		first_name,
		middle_name,
		last_name,
		type,
		status
	FROM users`

	rows, err := s.db.Leader.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []service.User
	for rows.Next() {
		var user service.User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.Type,
			&user.Status,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
