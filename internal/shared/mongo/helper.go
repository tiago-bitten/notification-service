package sharedMongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DefaultTimeout = 5 * time.Second

func WithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTimeout)
}

func FindOneDocument(
	ctx context.Context,
	collection *mongo.Collection,
	filter bson.M,
	result interface{},
) (bool, error) {
	err := collection.FindOne(ctx, filter).Decode(result)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return false, err
}
