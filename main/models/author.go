package models

type Author struct {
	BaseModel
	Name  string `gorm:"column:name;not null;uniqueIndex" json:"name"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books,omitempty"`
}

func (Author) TableName() string {
	return "authors"
}
