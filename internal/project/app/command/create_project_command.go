package command

import "github.com/tiago-bitten/notification-service/internal/project/domain/project"

type CreateProjectCommand struct {
	Name string
}

type CreateProjectHandler struct {
	projectRepo project.Repository
}

func NewCreateProjectHandler(projectRepo project.Repository) *CreateProjectHandler {
	return &CreateProjectHandler{
		projectRepo: projectRepo,
	}
}

func (h *CreateProjectHandler) Handle(cmd CreateProjectCommand) (string, error) {
	p := project.NewProject(cmd.Name)
	err := h.projectRepo.Save(p)

	return p.Key, err
}
