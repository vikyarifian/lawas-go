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

func ConnectDBVercel() {

	// dsn := os.Getenv("DB_USER_VERCEL") + ":" + os.Getenv("DB_PASS_VERCEL") + "@tcp(" + os.Getenv("DB_HOST_VERCEL") + ":" + os.Getenv("DB_PORT_VERCEL") + ")/" + os.Getenv("DB_NAME_VERCEL") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "sql12763547:xsy9uBDeMt@tcp(sql12.freesqldatabase.com:3306)/sql12763547?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "eternalb_lawas:Law@s(0-=@tcp(eternalbuana.id:3306)/eternalb_lawas?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "mysql_articletoy:489143e292dbb496707d8df847267d1c999e50b8@tcp(ji8k9.h.filess.io:3307)/mysql_articletoy?charset=utf8mb4&parseTime=True&loc=Local"
	// println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect Database MySQL  ("+os.Getenv("DB_HOST_VERCEL")+"/"+os.Getenv("DB_NAME_VERCEL")+")", err.Error())
		os.Exit(2)
	}

	log.Println("Connected To MySql Database Successfully (" + os.Getenv("DB_HOST_VERCEL") + "/" + os.Getenv("DB_NAME_VERCEL") + ")")
	MySql = db

}

// func ConnectDBVercel() {

// 	dsn := "sql12763547:xsy9uBDeMt@tcp(sql12.freesqldatabase.com:3306/sql12763547?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed To Connect Database MySQL", err.Error())
// 		os.Exit(2)
// 	}

// 	log.Println("Connected To MySql Database Successfully")
// 	MySql = db

// }
