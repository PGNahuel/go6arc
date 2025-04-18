package services

import (
	"context"
	"errors"

	"github.com/pgnahuel/go6arc/internal/domain"
	"github.com/pgnahuel/go6arc/internal/ports"
)

type BasicService struct {
	basicReeader    ports.BasicReader
	basicCreator    ports.BasicCreator
	basicEliminator ports.BasicEliminator
	basicRemover    ports.BasicRemover
}

func NewBasicService(basicReeader ports.BasicReader, basicCreator ports.BasicCreator, basicEliminator ports.BasicEliminator, basicRemover ports.BasicRemover) ports.BasicService {
	return &BasicService{
		basicReeader:    basicReeader,
		basicCreator:    basicCreator,
		basicEliminator: basicEliminator,
		basicRemover:    basicRemover,
	}
}

func (bs BasicService) Create(ctx context.Context, entity domain.BasicObject) error {
	return errors.New("not implemented")
}

func (bs BasicService) Read(ctx context.Context, id string) (domain.BasicObject, error) {
	return domain.BasicObject{}, errors.New("not implemented")
}

func (bs BasicService) Update(ctx context.Context, entity domain.BasicObject) error {
	return errors.New("not implemented")
}

func (bs BasicService) Delete(ctx context.Context, id string) error {
	return errors.New("not implemented")
}

func (bs BasicService) List(ctx context.Context) ([]domain.BasicObject, error) {
	return []domain.BasicObject{}, errors.New("not implemented")
}
