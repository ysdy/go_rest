package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ysdy/go_rest/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
    // MySQLとの接続。ユーザ名：gorm パスワード：password DB名：country
	db, err = gorm.Open("mysql", "{user:password}@({AWS RDS Address})/{DB NAME}?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
	
	autoMigration()
}

func autoMigration() {
   db.AutoMigrate(&entity.Pet{})
}
