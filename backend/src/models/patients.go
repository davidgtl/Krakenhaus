package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Patient struct {
	gorm.Model
	Name string `json:"name"`
	BirthDate *time.Time `json:"birth_date"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	MedicalRecord string `json:"medical_record"`
}