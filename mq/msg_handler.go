package mq

import (
	"flash-sale-backend/dao"
	"flash-sale-backend/model"
	"flash-sale-backend/utils"
)

func messageHandler(m []byte) {
	id := utils.StringToInt64(string(m))
	if id == 0 {
		return
	}

	stock := dao.LockStockByID(id)
	if stock < 0 {
		return
	}

	if stock == 0 {
		SyncStock(id)
	}

	data := new(model.SPUOrder)
	data.ActivityId = id
	data.OrderId = utils.GenerateID()
	if err := dao.CreateOrder(data); err != nil {
		SyncStock(id)
	}

}

func SyncStock(id int64) {
	// activity := dao.GetActivityByID(id)
	// if activity != nil {
	// 	dao.SetStock(id, activity.AvailStock)
	// }
}
