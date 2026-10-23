package repository

import (
	"context"

	"bitgroup-cs/internal/entity"
)

type ProductRepository interface {
	FetchAll(ctx context.Context) ([]*entity.Product, error)
	FindById(ctx context.Context, id int64) (*entity.Product, error)
	Create(ctx context.Context, data *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, data *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id int64) error
}
