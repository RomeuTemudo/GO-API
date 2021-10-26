package database

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

//connection
func Connect() {

	connection, err := gorm.Open(mysql.Open("root:rootroot@/eqs_schema"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")

	}

	DB = connection

}
