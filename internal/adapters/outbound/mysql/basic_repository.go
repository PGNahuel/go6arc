package mysql

import (
	"context"
	"database/sql"

	"github.com/pgnahuel/go6arc/internal/domain"
)

// Repo Mysql
type basicMySQLRepository struct {
	db *sql.DB
}

func NewBasicRepository() *basicMySQLRepository {
	return &basicMySQLRepository{
		db: nil,
	}

}

// Implementacion de los metodos
func (r *basicMySQLRepository) Create(ctx context.Context, entity domain.BasicObject) error {
	return nil
}

func (r *basicMySQLRepository) Read(ctx context.Context, id string) (domain.BasicObject, error) {
	return domain.BasicObject{}, nil
}

func (r *basicMySQLRepository) Update(ctx context.Context, entity domain.BasicObject) error {
	return nil
}

func (r *basicMySQLRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *basicMySQLRepository) List(ctx context.Context) ([]domain.BasicObject, error) {
	return []domain.BasicObject{}, nil
}
