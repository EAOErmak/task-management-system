package requests

import (
	"go-learn/main/models"
)

type BookCreateRequest struct {
	Title      string  `json:"title" binding:"required"`
	AuthorID   uint    `json:"author_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type BookResponse struct {
	ID         uint              `json:"id"`
	Title      string            `json:"title"`
	AuthorID   uint              `json:"author_id"`
	Author     *AuthorResponse   `json:"author,omitempty"`
	CategoryID uint              `json:"category_id"`
	Category   *CategoryResponse `json:"category,omitempty"`
	Price      float64           `json:"price"`
}

func NewBookResponse(book models.Book) BookResponse {
	response := BookResponse{
		ID:         book.ID,
		Title:      book.Title,
		AuthorID:   book.AuthorID,
		CategoryID: book.CategoryID,
		Price:      book.Price,
	}

	if book.Author != nil {
		author := NewAuthorResponse(*book.Author)
		response.Author = &author
	}

	if book.Category != nil {
		category := NewCategoryResponse(*book.Category)
		response.Category = &category
	}

	return response
}

func NewBookResponses(books []models.Book) []BookResponse {
	responses := make([]BookResponse, 0, len(books))
	for _, book := range books {
		responses = append(responses, NewBookResponse(book))
	}

	return responses
}
