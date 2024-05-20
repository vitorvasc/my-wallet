package domain

import "time"

type Transaction struct {
	ID        uint64
	UserID    uint64
	Type      string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
