package database

import (
	"flexusma.de/quashier-server/database/models"
	"flexusma.de/quashier-server/util"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	if util.Dev {
		db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database " + err.Error())
		}
		util.DBI = db
	} else {
		db, err := gorm.Open(mysql.Open(util.DBURL), &gorm.Config{})
		if err != nil {
			panic("failed to connect database " + err.Error())
		}
		util.DBI = db
	}
	println("Database connected, running migrations...")

	err := util.DBI.AutoMigrate(&models.User{}, &models.Item{}, &models.Purchase{})
	if err != nil {
		panic("Could not migrate models. This is unrecoverable, panicking: " + err.Error())
	}

	println("Migrations ran successful")

}
