package database

import (
	"fmt"
	"goLANG/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	connErr := godotenv.Load()
	if connErr != nil {
		log.Fatalf("Error connecting to .env file")
	}

	DBUSER := os.Getenv("DB_USERNAME")
	DBPASS	:= os.Getenv("DB_PASSWORD")
	DBPORT	:= os.Getenv("DB_PORT")
	DBHOST	:= os.Getenv("DB_HOST")
	DBNAME	:= os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("user=%s password=%s port=%s host=%s dbname=%s sslmode=disable", DBUSER, DBPASS, DBPORT, DBHOST, DBNAME)
	
	var gormErr error

	db, gormErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if gormErr != nil {
		panic(gormErr.Error())
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		panic(sqlDBErr.Error())
	}

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		panic(pingErr.Error())
	}

	fmt.Println("Success connect to database")
}

func Migrate() {
	db.AutoMigrate(
		&models.MasterUser{},
		&models.MasterDepartment{},
	)
}