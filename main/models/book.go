package models

type Book struct {
	BaseModel
	Title         string         `gorm:"column:title;not null" json:"title"`
	AuthorID      uint           `gorm:"column:author_id;not null;index" json:"author_id"`
	Author        *Author        `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	CategoryID    uint           `gorm:"column:category_id;not null;index" json:"category_id"`
	Category      *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Price         float64        `gorm:"column:price;not null" json:"price"`
	FavoriteBooks []FavoriteBook `gorm:"foreignKey:BookID" json:"-"`
}

func (Book) TableName() string {
	return "books"
}
