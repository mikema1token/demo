package cache

import (
	"context"
	"demo/db"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client

const UserListCacheKey = "user-list"

func init() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     "",
			Password: "",
			DB:       0,
		})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func SetUserListCache(userModelList []db.UserModel) error {
	cacheValue, err := json.Marshal(userModelList)
	if err != nil {
		return err
	}
	return redisClient.Set(context.Background(), UserListCacheKey, string(cacheValue), time.Hour).Err()
}

func DeleteUserListCache() error {
	return redisClient.Del(context.Background(), UserListCacheKey).Err()
}

func GetUserListCache() ([]db.UserModel, error) {
	cacheValue, err := redisClient.Get(context.Background(), UserListCacheKey).Result()
	if err != nil {
		return nil, err
	}
	var userModelList []db.UserModel
	err = json.Unmarshal([]byte(cacheValue), &userModelList)
	return userModelList, err
}
