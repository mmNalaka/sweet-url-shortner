package repository

import (
    "context"
	"sweet/internal/model"
)

type DoaminRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Doamin, error)
}

func NewDoaminRepository(repository *Repository) DoaminRepository {
	return &doaminRepository{
		Repository: repository,
	}
}

type doaminRepository struct {
	*Repository
}

func (r *doaminRepository) FirstById(ctx context.Context, id int64) (*model.Doamin, error) {
	var doamin model.Doamin
	// TODO: query db
	return &doamin, nil
}
