package models

import (
	"github.com/jinzhu/gorm"
)

type Medication struct {
	gorm.Model
	Name string `json:"name"`
	SideEffects string `json:"name"`
	Dosage int `json:"name"`
}