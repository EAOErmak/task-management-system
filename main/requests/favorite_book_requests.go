package requests

import (
	"go-learn/main/models"
)

type FavoriteBookCreateRequest struct {
	BookID uint `json:"book_id" binding:"required"`
}

type FavoriteBookResponse struct {
	ID     uint          `json:"id"`
	UserID uint          `json:"user_id"`
	BookID uint          `json:"book_id"`
	Book   *BookResponse `json:"book,omitempty"`
}

func NewFavoriteBookResponse(favorite models.FavoriteBook) FavoriteBookResponse {
	response := FavoriteBookResponse{
		ID:     favorite.ID,
		UserID: favorite.UserID,
		BookID: favorite.BookID,
	}

	if favorite.Book != nil {
		book := NewBookResponse(*favorite.Book)
		response.Book = &book
	}

	return response
}

func NewFavoriteBookResponses(favorites []models.FavoriteBook) []FavoriteBookResponse {
	responses := make([]FavoriteBookResponse, 0, len(favorites))
	for _, favorite := range favorites {
		responses = append(responses, NewFavoriteBookResponse(favorite))
	}

	return responses
}
