package requests

import (
	"go-learn/main/models"
)

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func NewCategoryResponses(categories []models.Category) []CategoryResponse {
	responses := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		responses = append(responses, NewCategoryResponse(category))
	}

	return responses
}
