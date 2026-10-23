package usecase

import (
	"context"

	"bitgroup-cs/internal/entity"
	appErrors "bitgroup-cs/internal/errors"
	"bitgroup-cs/internal/repository"
)

type IProductUseCase interface {
	FetchAll(ctx context.Context) ([]*entity.Product, error)
	FindById(ctx context.Context, id int64) (*entity.Product, error)
	CreateProduct(ctx context.Context, data *entity.Product) (*entity.Product, error)
	UpdateProductByID(ctx context.Context, data *entity.Product) (*entity.Product, error)
	DeleteProductByID(ctx context.Context, id int64) error
}

type ProductUseCase struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

func NewProductUseCase(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) IProductUseCase {
	return &ProductUseCase{
		productRepository:  productRepo,
		categoryRepository: categoryRepo,
	}
}

func (uc *ProductUseCase) FetchAll(ctx context.Context) ([]*entity.Product, error) {
	result, err := uc.productRepository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ProductUseCase) FindById(ctx context.Context, id int64) (*entity.Product, error) {
	result, err := uc.productRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, appErrors.New(appErrors.ErrNotFound, "Product not found")
	}
	return result, nil
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	if data == nil || data.Name == "" {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Product name is required")
	}
	if data.Price < 0 {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Product price cannot be negative")
	}
	if data.CategoryID <= 0 {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Valid Category ID is required")
	}

	category, err := uc.categoryRepository.FindById(ctx, data.CategoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Category does not exist")
	}

	result, err := uc.productRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ProductUseCase) UpdateProductByID(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	if data == nil || data.ID <= 0 {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Invalid product ID")
	}

	currentRecord, err := uc.productRepository.FindById(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if currentRecord == nil {
		return nil, appErrors.New(appErrors.ErrNotFound, "Product not found")
	}

	if data.CategoryID > 0 {
		category, err := uc.categoryRepository.FindById(ctx, data.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, appErrors.New(appErrors.ErrBadRequest, "Category does not exist")
		}
	} else {
		data.CategoryID = currentRecord.CategoryID
	}

	if data.Name == "" {
		data.Name = currentRecord.Name
	}
	if data.Price < 0 {
		return nil, appErrors.New(appErrors.ErrBadRequest, "Product price cannot be negative")
	}

	result, err := uc.productRepository.Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ProductUseCase) DeleteProductByID(ctx context.Context, id int64) error {
	currentRecord, err := uc.productRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if currentRecord == nil {
		return appErrors.New(appErrors.ErrNotFound, "Product not found")
	}

	if err := uc.productRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
