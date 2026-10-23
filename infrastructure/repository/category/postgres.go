package category

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

func NewPostgresRepository(dbConn *gorm.DB) repository.CategoryRepository {
	return &postgresRepository{
		dbConn: dbConn,
	}
}

func (repo *postgresRepository) FetchAll(ctx context.Context) ([]*entity.Category, error) {
	var categoryModels []CategoryModel
	err := repo.dbConn.WithContext(ctx).Order("id asc").Find(&categoryModels).Error
	if err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to fetch categories: %v", err))
	}
	return listCategoryModel(categoryModels).ToEntities(), nil
}

func (repo *postgresRepository) FindById(ctx context.Context, id int64) (*entity.Category, error) {
	var categoryModel CategoryModel
	err := repo.dbConn.WithContext(ctx).First(&categoryModel, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to find category by id: %v", err))
	}
	return categoryModel.ToEntity(), nil
}

func (repo *postgresRepository) Create(ctx context.Context, data *entity.Category) (*entity.Category, error) {
	model := CategoryModel{}.FromEntity(data)
	err := repo.dbConn.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to insert category: %v", err))
	}
	return model.ToEntity(), nil
}

func (repo *postgresRepository) Update(ctx context.Context, data *entity.Category) (*entity.Category, error) {
	model := CategoryModel{}.FromEntity(data)
	res := repo.dbConn.WithContext(ctx).Model(&CategoryModel{}).
		Where("id = ?", model.ID).
		Updates(map[string]interface{}{
			"name":       model.Name,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if res.Error != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to update category: %v", res.Error))
	}

	var updatedModel CategoryModel
	if err := repo.dbConn.WithContext(ctx).First(&updatedModel, "id = ?", model.ID).Error; err != nil {
		return nil, appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to retrieve updated category: %v", err))
	}

	return updatedModel.ToEntity(), nil
}

func (repo *postgresRepository) Delete(ctx context.Context, id int64) error {
	res := repo.dbConn.WithContext(ctx).Delete(&CategoryModel{}, "id = ?", id)
	if res.Error != nil {
		return appErrors.New(appErrors.ErrInternal, fmt.Sprintf("failed to delete category: %v", res.Error))
	}
	return nil
}
