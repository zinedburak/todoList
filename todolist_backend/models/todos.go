package models


type Todos struct {
	Id     uint   `json:"id"`
	Text   string `json:"content" gorm:"unique"`
}