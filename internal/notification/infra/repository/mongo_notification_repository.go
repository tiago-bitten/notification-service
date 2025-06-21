package repository

import (
	"github.com/tiago-bitten/notification-service/internal/notification/domain/notification"
	sharedMongo "github.com/tiago-bitten/notification-service/internal/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoNotificationRepository struct {
	collection *mongo.Collection
}

func NewMongoNotificationRepository(db *mongo.Database) *MongoNotificationRepository {
	return &MongoNotificationRepository{
		collection: db.Collection("notifications"),
	}
}

func (r *MongoNotificationRepository) Save(n *notification.Notification) error {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()
	_, err := r.collection.InsertOne(ctx, n)
	return err
}

func (r *MongoNotificationRepository) FindByID(ID string) (*notification.Notification, error) {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()

	var n notification.Notification
	found, err := sharedMongo.FindOneDocument(ctx, r.collection, bson.M{"_id": ID}, &n)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return &n, nil
}
