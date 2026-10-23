package product

import (
	"context"
	"errors"
	"fmt"

	"bitgroup-cs/internal/entity"
	appErrors "bitgroup-cs/internal/errors"
	"bitgroup-cs/internal/repository"
	"gorm.io/gorm"
)

type postgresRepository struct {
	dbConn *gorm.DB
}

func NewPostgresRepository(dbConn *gorm.DB) repository.ProductRepository {
	return &postgresRepository{
		dbConn: dbConn,
	}
}

func (repo *postgresRepository) FetchAll(ctx context.Context) ([]*entity.Product, error) {
	var productModels []ProductModel
	err := repo.dbConn.WithContext(ctx).Preload("Category").Order("id asc").Find(&productModels).Error
	if err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to fetch products: %v", err))
	}
	return listProductModel(productModels).ToEntities(), nil
}

func (repo *postgresRepository) FindById(ctx context.Context, id int64) (*entity.Product, error) {
	var productModel ProductModel
	err := repo.dbConn.WithContext(ctx).Preload("Category").First(&productModel, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to find product by id: %v", err))
	}
	return productModel.ToEntity(), nil
}

func (repo *postgresRepository) Create(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	model := ProductModel{}.FromEntity(data)
	err := repo.dbConn.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to insert product: %v", err))
	}

	// Preload category relation for the newly created product
	var createdModel ProductModel
	if err := repo.dbConn.WithContext(ctx).Preload("Category").First(&createdModel, "id = ?", model.ID).Error; err != nil {
		return model.ToEntity(), nil
	}

	return createdModel.ToEntity(), nil
}

func (repo *postgresRepository) Update(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	model := ProductModel{}.FromEntity(data)
	res := repo.dbConn.WithContext(ctx).Model(&ProductModel{}).
		Where("id = ?", model.ID).
		Updates(map[string]interface{}{
			"name":        model.Name,
			"price":       model.Price,
			"category_id": model.CategoryID,
			"updated_at":  gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if res.Error != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to update product: %v", res.Error))
	}

	var updatedModel ProductModel
	if err := repo.dbConn.WithContext(ctx).Preload("Category").First(&updatedModel, "id = ?", model.ID).Error; err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to retrieve updated product: %v", err))
	}

	return updatedModel.ToEntity(), nil
}

func (repo *postgresRepository) Delete(ctx context.Context, id int64) error {
	res := repo.dbConn.WithContext(ctx).Delete(&ProductModel{}, "id = ?", id)
	if res.Error != nil {
		return appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to delete product: %v", res.Error))
	}
	return nil
}
