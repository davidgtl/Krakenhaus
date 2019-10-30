package models

import (
	u "dslic/utils"
	"github.com/jinzhu/gorm"
	"reflect"
)

func Get(item interface{}) map[string] interface{}  {

	obj := reflect.ValueOf(item)
	mod := reflect.Indirect(obj).FieldByName("Model").Interface().(gorm.Model)

	item = GetDB().Where("id = ?", mod.ID).First(item)

	resp := u.Message(true, "success")
	resp["object"] = item
	return resp
}


func Create(item interface{}) map[string] interface{}  {

	GetDB().Create(item)

	resp := u.Message(true, "success")
	resp["object"] = item
	return resp
}

func Update(item interface{}) map[string] interface{} {

	GetDB().Save(item)

	resp := u.Message(true, "success")
	resp["object"] = item
	return resp
}

func List(items interface{}) map[string] interface{} {

	db.Find(items)

	resp := u.Message(true, "success")
	resp["object"] = items
	return resp
}

func Delete(item interface{}) {
	GetDB().Delete(item)
}
