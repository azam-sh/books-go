package models

type Category struct {
	ID   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name;unique"`
}

func (c Category) TableName() string {
	return "categories"
}
