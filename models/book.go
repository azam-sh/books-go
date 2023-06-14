package models

type Book struct {
	ID         int64    `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name       string   `json:"name" gorm:"column:name"`
	CategoryID int64    `json:"category_id" gorm:"column:category_id"`
	Category   Category `gorm:"foreignKey:category_id"`
}

type RequestBook struct {
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
}

func (b Book) TableName() string {
	return "books"
}
