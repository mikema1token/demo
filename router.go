package demo

import (
	"demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user:id", service.GetuserById)
	router.GET("/users", service.GetuserList)
	router.POST("/user", service.CreateUser)
	router.DELETE("/user:id", service.DeleteById)

}
