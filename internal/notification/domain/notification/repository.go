package notification

type Repository interface {
	Save(n *Notification) error
	FindByID(id string) (*Notification, error)
}
