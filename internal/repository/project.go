package repository

import (
    "context"
	"sweet/internal/model"
)

type ProjectRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Project, error)
}

func NewProjectRepository(repository *Repository) ProjectRepository {
	return &projectRepository{
		Repository: repository,
	}
}

type projectRepository struct {
	*Repository
}

func (r *projectRepository) FirstById(ctx context.Context, id int64) (*model.Project, error) {
	var project model.Project
	// TODO: query db
	return &project, nil
}
