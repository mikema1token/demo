package service

import (
	"demo/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
}

func NewUserFromUserModel(userModel db.UserModel) User {
	return User{
		Id:          int(userModel.ID),
		Name:        userModel.Name,
		CreatedTime: userModel.CreatedAt,
	}
}
func GetuserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err := db.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": NewUserFromUserModel(user)})
}

func GetuserList(c *gin.Context) {
	userModelList, err := db.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var userList []User
	for _, userModel := range userModelList {
		userList = append(userList, NewUserFromUserModel(userModel))
	}
	c.JSON(http.StatusOK, gin.H{"data": userList})
}

func CreateUser(c *gin.Context) {
	var newUser User

	// 从请求中解析 JSON 数据
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userModel := db.UserModel{
		Model: gorm.Model{ID: uint(newUser.Id)},
	}
	err := db.CreateUser(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": "ok"})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.DeleteUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
