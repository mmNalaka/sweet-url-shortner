package handler

import (
	"github.com/gin-gonic/gin"
	"sweet/internal/service"
)

type ProjectHandler struct {
	*Handler
	projectService service.ProjectService
}

func NewProjectHandler(handler *Handler, projectService service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		Handler:      handler,
		projectService: projectService,
	}
}

func (h *ProjectHandler) GetProject(ctx *gin.Context) {

}
