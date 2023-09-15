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

func StatPostOrderUserCountByDay(beginTime, endTime time.Time) (int, error) {
	var r struct {
		TotalUser int `json:"total_user"`
	}
	sql := `select count(*) as total_user from order where created_time between ? and ? group by user_id`
	err := db.Raw(sql, beginTime, endTime).Scan(&r).Error
	return r.TotalUser, err
}
