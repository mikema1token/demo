package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	gorm.Model
	Name string `json:"name"`
}

func GetUserById(id int) (UserModel, error) {
	var user UserModel
	err := db.First(&user, id).Error
	return user, err
}

func GetUserList() ([]UserModel, error) {
	var userList []UserModel
	err := db.Find(&userList).Error
	return userList, err
}

func CreateUser(user UserModel) error {
	return db.Create(user).Error
}

func DeleteUserById(id int) error {
	return db.Delete(&UserModel{}, id).Error
}
