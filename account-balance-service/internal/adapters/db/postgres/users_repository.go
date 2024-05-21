package postgres

import (
	"database/sql"
	"errors"

	"account-balance-service/internal/core/domain"
)

func (r *Repository) GetUserByID(userID uint64) (*domain.User, error) {
	stmt := "SELECT user_id FROM account_balance WHERE user_id = $1"
	row := r.db.QueryRow(stmt, userID)

	var user domain.User
	if err := row.Scan(&user.UserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}
