package repository

import (
	"context"
	"sweet/internal/model"
)

type LinkRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Link, error)
	FindFirstBYKey(ctx context.Context, key string) (*model.Link, error)
	Create(ctx context.Context, link *model.Link) error
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

func (r *linkRepository) FindFirstBYKey(ctx context.Context, key string) (*model.Link, error) {
	var link model.Link
	if err := r.DB(ctx).Where("key = ?", key).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *linkRepository) Create(ctx context.Context, link *model.Link) error {
	if err := r.DB(ctx).Create(link).Error; err != nil {
		return err
	}
	return nil
}
