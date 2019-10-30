package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MedicationPlan struct {
	gorm.Model
	StartDate *time.Time `json:"birth_date"`
	EndDate *time.Time `json:"birth_date"`
	Prescriptions []Prescription `json:"prescriptions"`
	Patient Patient `json:"patient"`
}