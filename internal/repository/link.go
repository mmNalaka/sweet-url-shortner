package repository

import (
    "context"
	"sweet/internal/model"
)

type LinkRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Link, error)
}

func NewLinkRepository(repository *Repository) LinkRepository {
	return &linkRepository{
		Repository: repository,
	}
}

type linkRepository struct {
	*Repository
}

func (r *linkRepository) FirstById(ctx context.Context, id int64) (*model.Link, error) {
	var link model.Link
	// TODO: query db
	return &link, nil
}
