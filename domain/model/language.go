package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Code string
	Name string
}
