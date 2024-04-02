package service

import (
    "context"
	"sweet/internal/model"
	"sweet/internal/repository"
)

type LinkService interface {
	GetLink(ctx context.Context, id int64) (*model.Link, error)
}

func NewLinkService(service *Service, linkRepository repository.LinkRepository) LinkService {
	return &linkService{
		Service:        service,
		linkRepository: linkRepository,
	}
}

type linkService struct {
	*Service
	linkRepository repository.LinkRepository
}

func (s *linkService) GetLink(ctx context.Context, id int64) (*model.Link, error) {
	return s.linkRepository.FirstById(ctx, id)
}
