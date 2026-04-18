package requests

import (
	"time"

	"go-learn/main/models"
)

type DictionaryItemCreateRequest struct {
	Type  string `json:"type" binding:"required"`
	Label string `json:"label" binding:"required"`
}

type DictionaryItemResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDictionaryItemResponse(item models.DictionaryItem) DictionaryItemResponse {
	return DictionaryItemResponse{
		ID:        item.ID,
		Type:      string(item.Type),
		Label:     item.Label,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func NewDictionaryItemResponses(items []models.DictionaryItem) []DictionaryItemResponse {
	responses := make([]DictionaryItemResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, NewDictionaryItemResponse(item))
	}

	return responses
}
