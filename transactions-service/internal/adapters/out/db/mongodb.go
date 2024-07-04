package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host       = "MONGODB_HOST"
	port       = "MONGODB_PORT"
	dbname     = "MONGODB_DATABASE"
	collection = "MONGODB_COLLECTION"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db.Collection(os.Getenv(collection)),
	}
}

func InitDB(ctx context.Context) *mongo.Database {
	conn := fmt.Sprintf("mongodb://%s:%s", os.Getenv(host), os.Getenv(port))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	db := client.Database(os.Getenv(dbname))
	return db
}
