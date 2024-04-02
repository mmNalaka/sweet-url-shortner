package service

import (
    "context"
	"sweet/internal/model"
	"sweet/internal/repository"
)

type DoaminService interface {
	GetDoamin(ctx context.Context, id int64) (*model.Doamin, error)
}

func NewDoaminService(service *Service, doaminRepository repository.DoaminRepository) DoaminService {
	return &doaminService{
		Service:        service,
		doaminRepository: doaminRepository,
	}
}

type doaminService struct {
	*Service
	doaminRepository repository.DoaminRepository
}

func (s *doaminService) GetDoamin(ctx context.Context, id int64) (*model.Doamin, error) {
	return s.doaminRepository.FirstById(ctx, id)
}
