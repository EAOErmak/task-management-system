package requests

import (
	"go-learn/main/models"
)

type AuthorCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type AuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewAuthorResponse(author models.Author) AuthorResponse {
	return AuthorResponse{
		ID:   author.ID,
		Name: author.Name,
	}
}

func NewAuthorResponses(authors []models.Author) []AuthorResponse {
	responses := make([]AuthorResponse, 0, len(authors))
	for _, author := range authors {
		responses = append(responses, NewAuthorResponse(author))
	}

	return responses
}
