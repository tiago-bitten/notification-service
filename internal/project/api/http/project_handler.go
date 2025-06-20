package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/project/app"
	"github.com/tiago-bitten/notification-service/internal/project/app/command"
	"net/http"
)

type ProjectHandler struct {
	app app.Application
}

func NewProject(app app.Application) *ProjectHandler {
	return &ProjectHandler{
		app: app,
	}
}

// CreateProject godoc
// @Summary Create a new project
// @Description Endpoint to create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Param request body command.CreateProjectCommand true "Project to create"
// @Success 201 {object} any
// @Failure 400 {object} any
// @Router /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := command.CreateProjectCommand{
		Name: req.Name,
	}

	key, err := h.app.Commands.CreateProject.Handle(cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, gin.H{"success": true, "key": key})
}

// FindAllProjects godoc
// @Summary Find all projects
// @Description Endpoint to find all projects
// @Tags Projects
// @Produce json
// @Success 200 {object} any
// @Failure 400 {object} any
// @Router /projects [get]
func (h *ProjectHandler) FindAllProjects(c *gin.Context) {
	views, err := h.app.Queries.FindAllProjects.Handle()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": views})
}
