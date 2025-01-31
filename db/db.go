package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var MySql *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect Database MySQL  ("+os.Getenv("DB_HOST")+"/"+os.Getenv("DB_NAME")+")", err.Error())
		os.Exit(2)
	}

	log.Println("Connected To MySql Database Successfully (" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME") + ")")
	MySql = db

}
