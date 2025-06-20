package repository

import (
	"context"
	"github.com/tiago-bitten/notification-service/internal/notification/domain/notification"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{Collection: collection}
}

func (r *MongoRepository) Save(n *notification.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, n)
	return err
}
