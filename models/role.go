package models

type Role struct {
	ID   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name;unique"`
}

func (r Role) TableName() string {
	return "roles"
}
