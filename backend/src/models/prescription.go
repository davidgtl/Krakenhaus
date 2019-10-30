package models

import (
	"github.com/jinzhu/gorm"
)

type Prescription struct {
	gorm.Model
	Medication Medication `json:"medication"`
	Intake int `json:"intake"`
}