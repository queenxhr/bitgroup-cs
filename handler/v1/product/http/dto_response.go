package product

import (
	"time"

	categoryHttp "bitgroup-cs/handler/v1/category/http"
	"bitgroup-cs/internal/entity"
)

type ProductObject struct {
	ID         int64                        `json:"id"`
	Name       string                       `json:"name"`
	Price      float64                      `json:"price"`
	CategoryID int64                        `json:"category_id"`
	Category   *categoryHttp.CategoryObject `json:"category,omitempty"`
	CreatedAt  time.Time                    `json:"created_at"`
	UpdatedAt  time.Time                    `json:"updated_at"`
}

func (ProductObject) ParseFromEntity(ent *entity.Product) *ProductObject {
	if ent == nil {
		return nil
	}
	obj := &ProductObject{
		ID:         ent.ID,
		Name:       ent.Name,
		Price:      ent.Price,
		CategoryID: ent.CategoryID,
		CreatedAt:  ent.CreatedAt,
		UpdatedAt:  ent.UpdatedAt,
	}
	if ent.Category != nil {
		obj.Category = categoryHttp.CategoryObject{}.ParseFromEntity(ent.Category)
	}
	return obj
}

type ProductResponse struct {
	Products []*ProductObject `json:"products"`
}

func (ProductResponse) ParseFromEntities(entities []*entity.Product) *ProductResponse {
	results := make([]*ProductObject, len(entities))
	for i, t := range entities {
		results[i] = ProductObject{}.ParseFromEntity(t)
	}
	return &ProductResponse{
		Products: results,
	}
}
