package repository

import (
	"context"

	"bitgroup-cs/internal/entity"
)

type CategoryRepository interface {
	FetchAll(ctx context.Context) ([]*entity.Category, error)
	FindById(ctx context.Context, id int64) (*entity.Category, error)
	Create(ctx context.Context, data *entity.Category) (*entity.Category, error)
	Update(ctx context.Context, data *entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, id int64) error
}
