package handler

import (
	"flash-sale-backend/dao"
	"flash-sale-backend/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type addActivityRequest struct {
	ActName    string  `json:"act_name"`
	SPUID      int64   `json:"spu_id"`
	SalePrice  float32 `json:"sale_price"`
	TotalTime  int64   `json:"total_time"`
	TotalStock int32   `json:"total_stock"`
	Prime      bool    `json:"prime"`
	Special    bool    `json:"special"`
}

func AddActivityHandler(c *gin.Context) {
	req := new(addActivityRequest)
	if err := c.BindJSON(req); err != nil {
		return
	}

	m := new(model.SPUActivity)
	m.ActName = req.ActName
	m.SpuId = req.SPUID
	m.SalePrice = req.SalePrice
	m.ActStatus = 0
	m.StartTime = time.Now()
	m.EndTime = m.StartTime.Add(time.Duration(req.TotalTime) * time.Second)
	m.TotalStock = req.TotalStock
	m.AvailStock = req.TotalStock
	m.LockStock = 0
	if req.Prime {
		m.Prime = 1
	}

	if req.Special {
		m.Special = 1
	}

	dao.CreateSPUActivity(m, req.TotalTime)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

type activityInfo struct {
	ActivityID int64    `json:"activity_id"`
	Title      string   `json:"title"`
	Tag        []string `json:"tag"`
	Image      string   `json:"image"`
	ShopName   string   `json:"shop_name"`
	EndTime    int64    `json:"end_time"`
	Stock      int32    `json:"stock"`
	Total      int32    `json:"total"`
	Price      float32  `json:"price"`
	PriceStr   string   `json:"price_str"`
}

func GetActivityInfoHandler(c *gin.Context) {
	resp := make([]*activityInfo, 0)
	activityList := dao.GetAllSPUActivity()
	spuIDs := make([]int64, 0)
	for _, activity := range activityList {
		spuIDs = append(spuIDs, activity.SpuId)
	}

	spuInfo := dao.GetSpuItemByID(spuIDs)
	for _, activity := range activityList {
		info := new(activityInfo)
		spu := spuInfo[activity.SpuId]
		if spu == nil {
			continue
		}
		info.ActivityID = activity.ID
		info.Title = spu.SPUName
		info.Tag = make([]string, 0, 2)
		if activity.Prime == 1 {
			info.Tag = append(info.Tag, "Prime")
		}

		if activity.Special == 1 {
			info.Tag = append(info.Tag, "Special")
		}

		info.Image = spu.SPUImage
		info.ShopName = "Lobo Shop"
		info.EndTime = activity.EndTime.Unix()
		info.Stock = activity.AvailStock
		info.Total = activity.TotalStock
		info.Price = activity.SalePrice
		info.PriceStr = fmt.Sprintf("$%.2f", info.Price)
		resp = append(resp, info)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
