package config

import (
	"dbo-be/entities"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	dataConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Username,
		AppConfig.Password,
		AppConfig.Host,
		AppConfig.Port,
		AppConfig.DatabaseName,
	)
	DB, err = gorm.Open(mysql.Open(dataConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return DB, nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&entities.User{},
		&entities.Order{},
	)

	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&entities.User{}, "Orders")
	DB.Migrator().CreateConstraint(&entities.User{}, "fk_users_orders")

	log.Println("Database Migration Completed...")

	return nil
}
