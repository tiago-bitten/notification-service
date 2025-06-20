package notification

func NewNotification(id string, title string, body string, type_ NotificationType, users []User) *Notification {
	return &Notification{
		ID:    id,
		Title: title,
		Body:  body,
		Type:  type_,
		Users: users,
	}
}

type Notification struct {
	ID    string
	Title string
	Body  string
	Type  NotificationType
	Users []User
}

type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypeSMS   NotificationType = "sms"
	NotificationTypePush  NotificationType = "push"
)

func NewUser(id string, name string, channel string) *User {
	return &User{
		ID:      id,
		Name:    name,
		Channel: channel,
	}
}

type User struct {
	ID      string
	Name    string
	Channel string
}
