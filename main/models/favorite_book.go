package models

type FavoriteBook struct {
	BaseModel
	UserID uint  `gorm:"column:user_id;not null;uniqueIndex:udx_favorite_books_user_book" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	BookID uint  `gorm:"column:book_id;not null;uniqueIndex:udx_favorite_books_user_book" json:"book_id"`
	Book   *Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;" json:"book,omitempty"`
}

func (FavoriteBook) TableName() string {
	return "favorite_books"
}
