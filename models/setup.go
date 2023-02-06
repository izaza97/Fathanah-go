package models

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connection user
// func ConnectDatabase() {
// 	db, err := gorm.Open("mysql", "fathanah:FathanahS3cr3t@tcp(157.245.207.179:5432)/fathanah")
// 	if err != nil {
// 		log.Println("Connection Failed to Open")
// 	} else {
// 		log.Println("Connection Established")
// 	}

// 	DB = db
// }
// func Connect() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(192.168.138.139:3306)/fathanah")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// func ConnectDatabase() {
// 	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/fathanah"))
// 	if err != nil {
// 		fmt.Println("Gagal koneksi database")
// 	}

// 	DB = db
// }

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("fathanah:FathanahS3cr3t@tcp(157.245.207.179:5432)/fathanah?charset=utf8&parseTime=True"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	DB = db
}

//192.168.138.139
//localhost
