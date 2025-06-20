package repository

import (
	"github.com/tiago-bitten/notification-service/internal/project/domain/project"
	"github.com/tiago-bitten/notification-service/internal/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProjectRepository struct {
	collection *mongo.Collection
}

func NewMongoProjectRepository(db *mongo.Database) *MongoProjectRepository {
	return &MongoProjectRepository{
		collection: db.Collection("projects"),
	}
}

func (r *MongoProjectRepository) Save(project *project.Project) error {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()
	_, err := r.collection.InsertOne(ctx, project)
	return err
}

func (r *MongoProjectRepository) FindAll() ([]project.Project, error) {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()

	var projects []project.Project
	cursor, err := r.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *MongoProjectRepository) FindByID(id string) (*project.Project, error) {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()

	var p project.Project
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	return &p, err
}

func (r *MongoProjectRepository) FindByName(name string) (*project.Project, error) {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()

	var p project.Project
	found, err := sharedMongo.FindOneDocument(ctx, r.collection, bson.M{"name": name}, &p)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return &p, nil
}

func (r *MongoProjectRepository) FindByKey(key string) (*project.Project, error) {
	ctx, cancel := sharedMongo.WithTimeout()
	defer cancel()

	var p project.Project
	found, err := sharedMongo.FindOneDocument(ctx, r.collection, bson.M{"key": key}, &p)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return &p, nil
}
