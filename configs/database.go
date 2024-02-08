package configs

import (
	"fmt"
	"os"
	"test_dealls/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Database()  {
	fmt.Println("TEST CONFIG")

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchmea := os.Getenv("DB_SCHEMA")


	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbSchmea + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
        PrepareStmt:            true,
	})

	if err != nil {
        panic(err)
    }

	DB = db

    db.AutoMigrate(
		models.Swipe{},
		models.User{},
	)

	fmt.Println("Connect Database " + dbConnection)

}