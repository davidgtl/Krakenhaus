package models

import (
	u "dslic/utils"
	. "dslic/validators"
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

func IsWellFormed(p *Patient) bool{
	return NonEmptyStr(p.Name, p.Gender, p.Address, p.MedicalRecord)
}

func (p *Patient) Create() map[string] interface{} {

	if !IsWellFormed(p) {
		return u.Message(false, "Empty field")
	}

	GetDB().Create(p)

	resp := u.Message(true, "success")
	resp["patient"] = p
	return resp
}

func (p *Patient) Update() map[string] interface{} {

	if !IsWellFormed(p) {
		return u.Message(false, "Empty field")
	}

	GetDB().Save(p)

	resp := u.Message(true, "success")
	resp["patient"] = p
	return resp
}
/*
func GetContacts(user uint) []*Patient {

	contacts := make([]*Patient, 0)
	err := GetDB().Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}*/
