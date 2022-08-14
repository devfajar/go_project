package models

import (
	"go-bookstore/pkg/config"

	"github.com/devfajar/go-bookstore/pkg/config"
)

var db *gorm.db

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
