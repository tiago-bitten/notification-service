package project

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strings"
	"time"
)

type Project struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Key      string             `json:"key" bson:"key"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
}

func NewProject(name string) *Project {
	return &Project{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Key:      generateKey(name),
		CreateAt: time.Now(),
	}
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateKey(projectName string) string {
	var b strings.Builder
	for i := 0; i < 15; i++ {
		b.WriteByte('0' + byte(rnd.Intn(10)))
	}
	return fmt.Sprintf("%s%s", projectName, b.String())
}
