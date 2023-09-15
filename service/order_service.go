package service

import (
	"demo/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Order struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	OrderSn     string    `json:"order_sn"`
	CreatedTime time.Time `json:"created_time"`
}

func NewOrderFromOrderModel(orderModel db.OrderModel) Order {
	return Order{
		Id:          int(orderModel.ID),
		UserId:      orderModel.UserId,
		OrderSn:     orderModel.OrderSn,
		CreatedTime: orderModel.CreatedAt,
	}
}
func GetOrderList(c *gin.Context) {
	type req struct {
		UserId         int    `json:"user_id"`
		OrderSn        string `json:"order_sn"`
		BeginTimestamp int    `json:"begin_timestamp"`
		EndTimestamp   int    `json:"end_timestamp"`
	}
	var r req
	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderModelList, err := db.GetOrderListByUserIdOrOrderSNOrCreatedTime(r.UserId, r.OrderSn, time.Unix(0, int64(r.BeginTimestamp)), time.Unix(0, int64(r.EndTimestamp)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var orderList []Order
	for _, orderModel := range orderModelList {
		orderList = append(orderList, NewOrderFromOrderModel(orderModel))
	}
	c.JSON(http.StatusOK, gin.H{"data": orderList})

}
