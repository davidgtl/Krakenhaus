package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dburl := os.Getenv("DATABASE_URL")

	dburlmatch := regexp.MustCompile(`(.*?):\/\/(.*?):(.*?)@(.*?):(.*?)\/(.*)`)
	match := dburlmatch.FindStringSubmatch(dburl)

	username := match[2]
	password := match[3]
	dbName := match[6]
	dbHost := match[4]


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	if false {
		//db.DropTable(&Account{})
		//db.DropTable(&Caregiver{})
		//db.DropTable(&MedicationPlan{})
		//db.DropTable(&Medication{})
		//db.DropTable(&Patient{})
		//db.DropTable(&Prescription{})
	}
	db.Debug().AutoMigrate(&Account{}, &Patient{}, &Medication{}, &MedicationPlan{}, &Caregiver{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
