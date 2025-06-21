package notification

func NewUser(id string, name string, channel string) *User {
	return &User{
		ID:      id,
		Name:    name,
		Channel: channel,
	}
}

type User struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Channel string `json:"channel" bson:"channel"`
}
