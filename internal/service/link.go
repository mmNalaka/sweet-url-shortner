package service

import (
	"context"
	v1 "sweet/api/v1"
	"sweet/internal/model"
	"sweet/internal/repository"
)

type LinkService interface {
	GetLink(ctx context.Context, id int64) (*model.Link, error)
	CreateLink(ctx context.Context, req *v1.CreateLinkRequest) (*v1.CreateLinkResponseData, error)
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

func (s *linkService) CreateLink(ctx context.Context, req *v1.CreateLinkRequest) (*v1.CreateLinkResponseData, error) {
	// Generate link KEY
	key, err := s.sid.GenString()
	if err != nil {
		return nil, err
	}

	link := &model.Link{
		Key: key,
		Url: req.Url,
	}

	if err = s.linkRepository.Create(ctx, link); err != nil {
		return nil, err
	}

	return &v1.CreateLinkResponseData{
		ShortLink: link.Key,
		Url:       link.Url,
	}, nil
}
