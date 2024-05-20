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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS account_balance (
		    user_id SERIAL PRIMARY KEY,
		    balance NUMERIC NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	return db
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (p PostgresRepository) GetAccountBalance(userID uint64) (float64, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) UpdateAccountBalance(userID uint64, amount float64) error {
	// TODO implement me
	panic("implement me")
}
