package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title		string `json:"title" form:"title"`
	Writer 		string `json:"writer" form:"writer"`
	Publisher 	string `json:"publisher" form:"publisher"`

}