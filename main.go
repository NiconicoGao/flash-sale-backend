package main

import (
	"flash-sale-backend/dao"
	"flash-sale-backend/middleware"
	"flash-sale-backend/mq"
	"flash-sale-backend/router"
	"flash-sale-backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	myConfig, err := utils.ReadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = utils.InitSnowflake()
	if err != nil {
		panic(err)
	}

	err = dao.DBInit(&myConfig.MysqlConfig)
	if err != nil {
		panic(err)
	}

	err = dao.RedisInit(&myConfig.RedisConfg)
	if err != nil {
		panic(err)
	}

	// err = kafka.KafkaInit(&myConfig.KafkaConfig)
	// if err != nil {
	// 	panic(err)
	// }

	err = mq.InitActiveMQ(&myConfig.ActiveMQConfig)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	address := fmt.Sprintf("%v:%v", myConfig.ServerConfig.IP, myConfig.ServerConfig.Port)
	fmt.Printf("Server Start at %v\n", address)
	r := gin.Default()
	r.Use(middleware.ConfigMiddleware(myConfig))
	router.HandleRouter(r)
	r.Run(address)
}
