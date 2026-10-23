package category

import (
	"time"

	"bitgroup-cs/internal/entity"
)

type CategoryObject struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (CategoryObject) ParseFromEntity(ent *entity.Category) *CategoryObject {
	if ent == nil {
		return nil
	}
	return &CategoryObject{
		ID:        ent.ID,
		Name:      ent.Name,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

type CategoryResponse struct {
	Categories []*CategoryObject `json:"categories"`
}

func (CategoryResponse) ParseFromEntities(entities []*entity.Category) *CategoryResponse {
	results := make([]*CategoryObject, len(entities))
	for i, t := range entities {
		results[i] = CategoryObject{}.ParseFromEntity(t)
	}
	return &CategoryResponse{
		Categories: results,
	}
}
