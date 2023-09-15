package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderModel struct {
	gorm.Model
	UserId  int    `json:"user_id"`
	OrderSn string `json:"order_sn"`
}

func GetOrderListByUserIdOrOrderSNOrCreatedTime(userId int, orderSn string, beginTime, endTime time.Time) ([]OrderModel, error) {
	var orderList []OrderModel
	err := db.Where("user_id = ? or order_sn = ? or created_time between ? and ? ", userId, orderSn, beginTime, endTime).Find(&orderList).Error
	return orderList, err
}
