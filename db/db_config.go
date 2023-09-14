package db

import "github.com/jinzhu/gorm"

var db gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3006)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&UserModel{})
}
