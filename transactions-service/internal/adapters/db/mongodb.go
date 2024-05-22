package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host       = "mongodb"
	port       = 27017
	dbname     = "transactions"
	collection = "transactions"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db.Collection(collection),
	}
}

func InitDB(ctx context.Context) *mongo.Database {
	conn := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	db := client.Database(dbname)
	return db
}
