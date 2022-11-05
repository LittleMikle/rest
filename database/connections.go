package database

import (
	"fmt"
	"github.com/LittleMikle/rest.git/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	//"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=123 dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connection success")
	}
	db.AutoMigrate(&models.Book{})
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = Init()
		sleep := time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			fmt.Printf("database is unavaible. wait for %d sec. \n", sleep)
			time.Sleep(sleep * time.Second)
			DB = Init()
		}
	}
	return DB
}
