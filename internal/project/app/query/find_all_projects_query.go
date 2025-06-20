package query

import "github.com/tiago-bitten/notification-service/internal/project/domain/project"

type projectView struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func newProjectView(p *project.Project) projectView {
	return projectView{
		Name: p.Name,
		Key:  p.Key,
	}
}

type FindAllProjectsQuery struct {
	projectRepo project.Repository
}

func NewFindAllProjectsQuery(projectRepo project.Repository) *FindAllProjectsQuery {
	return &FindAllProjectsQuery{
		projectRepo: projectRepo,
	}
}

func (q *FindAllProjectsQuery) Handle() ([]projectView, error) {
	projects, err := q.projectRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var views []projectView
	for _, p := range projects {
		views = append(views, newProjectView(&p))
	}

	return views, nil
}
