package models

import (
	"log"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("fathanah:FathanahS3cr3t@tcp(157.245.207.179:5432)/fathanah?charset=utf8&parseTime=True"))
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	DB = db
}

//192.168.138.139
//localhost
