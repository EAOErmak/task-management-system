package handlers

type namedEntityRequest struct {
	Name string `json:"name" binding:"required"`
}

type bookRequest struct {
	Title      string  `json:"title" binding:"required"`
	AuthorID   uint    `json:"author_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type favoriteBookRequest struct {
	BookID uint `json:"book_id" binding:"required"`
}
