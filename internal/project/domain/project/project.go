package project

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Project struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Key       string             `json:"key" bson:"key"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func NewProject(name string) *Project {
	return &Project{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Key:       generateKey(name),
		CreatedAt: time.Now(),
	}
}

func generateKey(projectName string) string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	return fmt.Sprintf("%s%s", projectName, id)
}
