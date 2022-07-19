package handler

import (
	"flash-sale-backend/dao"
	"flash-sale-backend/mq"
	"flash-sale-backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PlaceOrderResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var errorMsg = map[int]string{
	401: "Invalid input ID",
	402: "Activity haven't start",
	403: "Activity Ended",
	404: "Activity Not Found",
	405: "Item Sold Out",
	0:   "",
}

func response(code int) *PlaceOrderResp {
	resp := new(PlaceOrderResp)
	resp.Code = code
	resp.Message = errorMsg[code]
	return resp
}

func PlaceOrderHandler(c *gin.Context) {
	str := c.Query("id")
	id := utils.StringToInt64(str)

	if id == 0 {
		c.JSON(http.StatusOK, response(401))
		return
	}

	endTime := dao.GetActivityEndTimeByID(id)
	nowTime := time.Now().Unix()
	if nowTime > endTime {
		c.JSON(http.StatusOK, response(403))
		return
	}

	lock := dao.CheckAvalStock(id)
	if !lock {
		c.JSON(http.StatusOK, response(405))
		return
	}

	go mq.ProduceActiveMQ([]byte(fmt.Sprintf("%v", id)))
	c.JSON(http.StatusOK, response(0))
}
