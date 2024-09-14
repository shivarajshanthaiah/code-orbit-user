package db

import (
	"fmt"
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_user/config"
	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection todatabase failed: ", err)
	}

	err = DB.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Printf("error while migrating %v \n", err.Error())
		return nil
	}
	return DB
}
