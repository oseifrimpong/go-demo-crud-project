package config

import (
	"fmt"

	"github.com/obed/demo-project/models"

	"github.com/jinzhu/gorm"
	// Not needed for import.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Object used to access database tables and executing queries
var Database *gorm.DB
var user models.User

func init() {
	var err error
	Database, err = gorm.Open("postgres", "host=localhost port=5432 user=oseifrimpong dbname=godemo_db password=Frimpong.com20 sslmode=disable")

	if err != nil {
		panic(err)
	}

	// set this to 'true' to see sql logs
	Database.LogMode(true)
	Database.AutoMigrate(&user)

	fmt.Println("Database connection successful.")
}
