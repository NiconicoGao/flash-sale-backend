package dao

import (
	"flash-sale-backend/utils"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB
var redisConn *redis.Client

func DBInit(config *utils.MysqlConfig) error {
	dsn := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb3&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, config.Username, config.Password, config.IP, config.Port, config.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	conn = db
	return nil
}

func RedisInit(config *utils.RedisConfig) error {
	addr := fmt.Sprintf("%v:%v", config.IP, config.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}
