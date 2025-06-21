package notification

import "go.mongodb.org/mongo-driver/bson/primitive"

func NewNotification(
	title string,
	body string,
	type_ Type,
	projectKey string,
	users []User) *Notification {
	return &Notification{
		Title:      title,
		Body:       body,
		Type:       type_,
		ProjectKey: projectKey,
		Users:      users,
	}
}

type Notification struct {
	ID         primitive.ObjectID `json:"id" bson:"_id, omitempty "`
	Title      string             `json:"title" bson:"title"`
	Body       string             `json:"body" bson:"body"`
	Type       Type               `json:"type" bson:"type"`
	ProjectKey string             `json:"project_key" bson:"project_key"`
	Users      []User             `json:"users" bson:"users"`
}

type Type string

const (
	NotificationTypeEmail Type = "email"
	NotificationTypeSMS   Type = "sms"
	NotificationTypePush  Type = "push"
)
