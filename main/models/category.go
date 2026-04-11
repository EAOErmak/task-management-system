package models

type Category struct {
	BaseModel
	Name  string `gorm:"column:name;not null;uniqueIndex" json:"name"`
	Books []Book `gorm:"foreignKey:CategoryID" json:"books,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
