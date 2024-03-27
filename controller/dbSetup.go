package controller

import (
	"errors"
	"fmt"
	"gogin-basics/models"

	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrNotFound     = errors.New("record not found")
	ErrNotConnected = errors.New("connection to db not initialized")
)

var db *gorm.DB

func InitSampleDb(user string, pw string) error {

	cfgGorm := gormmysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(localhost:3306)/my_users_db?charset=utf8mb4&parseTime=True&loc=Local", user, pw),
	}

	var err error
	db, err = gorm.Open(gormmysql.Open(cfgGorm.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("Error in DB connection")
		return err
	}

	db.Table("user_profile").Model(&models.UserProfileEntity{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error in DB connection")
		return err
	}

	
	fmt.Println(sqlDB.Ping())

	return nil
}
