package models

import (
	"github.com/jinzhu/gorm"
)

type Medication struct {
	gorm.Model
	Name        string `json:"name"`
	SideEffects string `json:"side_effects"`
	Dosage      int    `json:"dosage"`
}
