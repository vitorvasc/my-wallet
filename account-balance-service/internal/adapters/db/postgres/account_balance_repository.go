package postgres

import "fmt"

func (r *Repository) GetAccountBalance(userID uint64) (float64, error) {
	stmt := "SELECT balance FROM account_balance WHERE user_id = $1"
	row := r.db.QueryRow(stmt, userID)

	var balance float64
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}

func (r *Repository) Debit(userID uint64, amount float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var balance float64
	err = tx.QueryRow("SELECT balance FROM account_balance WHERE user_id = $1 FOR UPDATE", userID).Scan(&balance)
	if err != nil {
		return err
	}

	if balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	_, err = tx.Exec("UPDATE account_balance SET balance = $1 WHERE user_id = $2", amount, userID)
	return err
}

func (r *Repository) Credit(userID uint64, amount float64) error {
	_, err := r.db.Exec("UPDATE account_balance SET balance = $1 WHERE user_id = $2", amount, userID)
	return err
}
