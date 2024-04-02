package repository

import (
    "context"
	"sweet/internal/model"
)

type TagRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Tag, error)
}

func NewTagRepository(repository *Repository) TagRepository {
	return &tagRepository{
		Repository: repository,
	}
}

type tagRepository struct {
	*Repository
}

func (r *tagRepository) FirstById(ctx context.Context, id int64) (*model.Tag, error) {
	var tag model.Tag
	// TODO: query db
	return &tag, nil
}
