package ports

import (
	"context"

	"github.com/pgnahuel/go6arc/internal/domain"
)

type BasicCreator interface {
	Create(ctx context.Context, entity domain.BasicObject) error
}

type BasicReader interface {
	Read(ctx context.Context, id string) (domain.BasicObject, error)
	List(ctx context.Context) ([]domain.BasicObject, error)
}

type BasicRemover interface {
	Delete(ctx context.Context, id string) error
}

type BasicEliminator interface {
	Update(ctx context.Context, entity domain.BasicObject) error
}
