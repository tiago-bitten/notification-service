package command

import (
	"github.com/tiago-bitten/notification-service/internal/project/domain/project"
)

type CreateProjectCommand struct {
	Name string `json:"name"`
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
	p, err := h.projectRepo.FindByName(cmd.Name)
	if err != nil {
		return "", err
	}

	if p != nil {
		return "", project.ErrProjectNameInUse
	}

	newProject := project.NewProject(cmd.Name)
	err = h.projectRepo.Save(newProject)

	return newProject.Key, err
}
