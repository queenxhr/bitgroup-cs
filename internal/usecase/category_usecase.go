package usecase

import (
	"context"

	"bitgroup-cs/internal/entity"
	appErrors "bitgroup-cs/internal/errors"
	"bitgroup-cs/internal/repository"
)

type ICategoryUseCase interface {
	FetchAll(ctx context.Context) ([]*entity.Category, error)
	FindById(ctx context.Context, id int64) (*entity.Category, error)
	CreateCategory(ctx context.Context, data *entity.Category) (*entity.Category, error)
	UpdateCategoryByID(ctx context.Context, data *entity.Category) (*entity.Category, error)
	DeleteCategoryByID(ctx context.Context, id int64) error
}

type CategoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepository) ICategoryUseCase {
	return &CategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *CategoryUseCase) FetchAll(ctx context.Context) ([]*entity.Category, error) {
	result, err := uc.categoryRepository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *CategoryUseCase) FindById(ctx context.Context, id int64) (*entity.Category, error) {
	result, err := uc.categoryRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, appErrors.New(appErrors.ErrNotFound, "Category not found")
	}
	return result, nil
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, data *entity.Category) (*entity.Category, error) {
	if data == nil || data.Name == "" {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Category name is required")
	}

	result, err := uc.categoryRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *CategoryUseCase) UpdateCategoryByID(ctx context.Context, data *entity.Category) (*entity.Category, error) {
	if data == nil || data.ID <= 0 {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Invalid category ID")
	}

	currentRecord, err := uc.categoryRepository.FindById(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if currentRecord == nil {
		return nil, appErrors.New(appErrors.ErrNotFound, "Category not found")
	}

	result, err := uc.categoryRepository.Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *CategoryUseCase) DeleteCategoryByID(ctx context.Context, id int64) error {
	currentRecord, err := uc.categoryRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if currentRecord == nil {
		return appErrors.New(appErrors.ErrNotFound, "Category not found")
	}

	if err := uc.categoryRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
