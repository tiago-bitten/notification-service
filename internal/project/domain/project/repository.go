package project

type Repository interface {
	Save(p *Project) error
	FindAll() ([]Project, error)
	FindByID(id string) (*Project, error)
	FindByName(name string) (*Project, error)
	FindByKey(key string) (*Project, error)
}
