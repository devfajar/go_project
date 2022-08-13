package models

var db *gorm.db

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
}
