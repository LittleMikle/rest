package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
