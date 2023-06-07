package models

type Book struct {
	ID         int64    `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name       string   `json:"name" gorm:"column:name"`
	AccessID   int64    `json:"accessId" gorm:"column:access_id"`
	Access     Access   `gorm:"foreignKey:AccessID"`
	CategoryID int64    `json:"categoryId" gorm:"column:category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type RequestBook struct {
	Name       string `json:"name"`
	AccessID   int64  `json:"accessId"`
	CategoryID int64  `json:"categoryId"`
}

func (b Book) TableName() string {
	return "books"
}
