package handler

import (
	"flash-sale-backend/dao"
	"flash-sale-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addItemRequest struct {
	Name  string
	Price float32
	Desc  string
	Image string
}

func AddItemHandler(c *gin.Context) {
	req := new(addItemRequest)
	if err := c.BindJSON(req); err != nil {
		return
	}

	spuItem := new(model.SPUItem)
	spuItem.SPUDesc = req.Desc
	spuItem.SPUName = req.Name
	spuItem.SPUImage = req.Image
	spuItem.SPUPrice = req.Price
	dao.CreateSPUItem(spuItem)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type itemInfo struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	PriceStr string  `json:"price_str"`
	Desc     string  `json:"desc"`
	Image    string  `json:"image"`
}

func GetAllItemHandler(c *gin.Context) {
	itemList := dao.GetAllSPUInfo()
	infoList := make([]*itemInfo, 0)
	for _, item := range itemList {
		infoList = append(infoList, &itemInfo{
			ID:       item.ID,
			Name:     item.SPUName,
			Price:    item.SPUPrice,
			PriceStr: fmt.Sprintf("$%.2f", item.SPUPrice),
			Desc:     item.SPUDesc,
			Image:    item.SPUImage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": infoList,
	})

}
