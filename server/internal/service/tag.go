package service

import (
    "context"
	"sweet/internal/model"
	"sweet/internal/repository"
)

type TagService interface {
	GetTag(ctx context.Context, id int64) (*model.Tag, error)
}

func NewTagService(service *Service, tagRepository repository.TagRepository) TagService {
	return &tagService{
		Service:        service,
		tagRepository: tagRepository,
	}
}

type tagService struct {
	*Service
	tagRepository repository.TagRepository
}

func (s *tagService) GetTag(ctx context.Context, id int64) (*model.Tag, error) {
	return s.tagRepository.FirstById(ctx, id)
}
