package db

import (
	"context"
	"log"

	"transactions-service/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) CreateTransaction(transaction *domain.Transaction) error {
	res, err := r.collection.InsertOne(context.Background(), transaction)
	if err != nil {
		log.Printf("[TransactionRepository_CreateTransaction] Error creating transaction: %v", err)
		return domain.ErrCreatingTransaction
	}

	transaction.ID = res.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("[repository] Transaction created with ID: %s", res.InsertedID)
	return nil
}

func (r *MongoRepository) UpdateTransaction(transaction *domain.Transaction) error {
	res, err := r.collection.UpdateOne(context.Background(), bson.M{"id": transaction.ID}, bson.M{"$set": transaction})
	if err != nil {
		return domain.ErrUpdatingTransaction
	}
	log.Printf("[repository] Transaction updated with ID: %s", res.UpsertedID)
	return nil
}

func (r *MongoRepository) GetTransactionByID(transactionID uint64) (*domain.Transaction, error) {
	res := r.collection.FindOne(context.Background(), bson.M{"id": transactionID})
	transaction := new(domain.Transaction)
	err := res.Decode(transaction)
	if err != nil {
		return nil, domain.ErrParsingTransaction
	}
	return transaction, nil
}

func (r *MongoRepository) GetTransactionsByUserID(userID uint64) ([]*domain.Transaction, error) {
	res, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, domain.ErrObtainingTransaction
	}
	transactions := make([]*domain.Transaction, 0)
	err = res.All(context.Background(), &transactions)
	if err != nil {
		return nil, domain.ErrParsingTransaction
	}
	return transactions, nil
}
