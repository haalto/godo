package data

import (
	"github.com/haalto/godo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dbConf := config.GetConfig().Database
	db, err := gorm.Open(postgres.Open(dbConf.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
