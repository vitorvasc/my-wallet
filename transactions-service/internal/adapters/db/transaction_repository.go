package db

import "transactions-service/internal/core/domain"

func (r *MongoRepository) CreateTransaction(transaction *domain.Transaction) error {
	// TODO implement me
	panic("implement")
}

func (r *MongoRepository) UpdateTransaction(transaction *domain.Transaction) error {
	// TODO implement me
	panic("implement")
}

func (r *MongoRepository) GetTransactionByID(transactionID uint64) (*domain.Transaction, error) {
	// TODO implement me
	panic("implement me")
}

func (r *MongoRepository) GetTransactionsByUserID(userID uint64) ([]*domain.Transaction, error) {
	// TODO implement me
	panic("implement me")
}
