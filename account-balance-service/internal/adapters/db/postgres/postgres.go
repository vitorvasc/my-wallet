package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "PG_HOST"
	port     = "PG_PORT"
	user     = "PG_USERNAME"
	password = "PG_PASSWORD"
	dbname   = "PG_DBNAME"
)

type Repository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func InitDB() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv(host), os.Getenv(port), os.Getenv(user), os.Getenv(password), os.Getenv(dbname))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error pinging db: %v", err)
	}

	return db
}
