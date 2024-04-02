package service

import (
    "context"
	"sweet/internal/model"
	"sweet/internal/repository"
)

type ProjectService interface {
	GetProject(ctx context.Context, id int64) (*model.Project, error)
}

func NewProjectService(service *Service, projectRepository repository.ProjectRepository) ProjectService {
	return &projectService{
		Service:        service,
		projectRepository: projectRepository,
	}
}

type projectService struct {
	*Service
	projectRepository repository.ProjectRepository
}

func (s *projectService) GetProject(ctx context.Context, id int64) (*model.Project, error) {
	return s.projectRepository.FirstById(ctx, id)
}
