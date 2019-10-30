package models

import (
	//"errors"
	u "dslic/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

func Get(item interface{}) error {

	obj := reflect.ValueOf(item)

	/*objelem := obj.Elem()
	for i := 0; i < objelem.NumField(); i++ {
		fmt.Println(objelem.Type().Field(i).Name)
		fmt.Println(objelem.Type().Field(i))
	}*/

	mod := reflect.Indirect(obj).FieldByName("Model").Interface().(gorm.Model)

	fmt.Print(mod.ID)

	//fmt.Print(reflect.ValueOf(intr).FieldByName("ID").Interface())
	//val := reflect.ValueOf(&item).Elem()

	id := mod.ID
	return GetDB().Where("id = ?", id).First(item).Error

	//return errors.New("argument is not a gorm model")
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

	db.Find(&items)

	resp := u.Message(true, "success")
	resp["object"] = items
	return resp
}

func Delete(item interface{}) {
	GetDB().Delete(item)
}
