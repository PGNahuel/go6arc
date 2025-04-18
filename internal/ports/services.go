package ports

import (
	"context"

	"github.com/pgnahuel/go6arc/internal/domain"
)

type BasicService interface {
	Create(ctx context.Context, entity domain.BasicObject) error
	Read(ctx context.Context, id string) (domain.BasicObject, error)
	Update(ctx context.Context, entity domain.BasicObject) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.BasicObject, error)
}
