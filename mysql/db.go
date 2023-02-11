package mysql

import (
	"golang-webapi/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pswd := os.Getenv("DB_PSWD")
	dbname := os.Getenv("DB_NAME")

	dsn := user + ":" + pswd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("MySQL server is connected...")
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{})

	log.Println("User table migration completed !!!")
}
