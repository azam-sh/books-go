package models

type Access struct {
	ID   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name;unique"` 
}

func (a Access) TableName() string {
	return "accesses"
}
