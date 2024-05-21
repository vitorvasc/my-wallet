package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "account_balance"
)

type PostgresRepository struct {
	db *sql.DB
}

func InitDB() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error pinging db: %v", err)
	}

	return db
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (r PostgresRepository) GetAccountBalance(userID uint64) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM account_balance WHERE user_id = $1", userID).Scan(&balance)
	return balance, err
}

func (r PostgresRepository) Debit(userID uint64, amount float64) error {
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

func (r PostgresRepository) Accredit(userID uint64, amount float64) error {
	_, err := r.db.Exec("UPDATE account_balance SET balance = $1 WHERE user_id = $2", amount, userID)
	return err
}
