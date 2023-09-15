package main

import (
	"demo/db"
	"demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/user:id", service.GetuserById)
	router.GET("/users", service.GetuserList)
	router.POST("/user", service.CreateUser)
	router.DELETE("/user:id", service.DeleteById)
	router.GET("/order/list", service.GetOrderList)
	go LoopTask()
	router.Run()
}

func LoopTask() {
	ticker := time.NewTicker(time.Hour * 24)
	for _ = range ticker.C {
		totalUser, err := db.StatPostOrderUserCountByDay(time.Now().AddDate(0, 0, -1), time.Now())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("过去一天下单的用户有:", totalUser)
		}
	}
}
