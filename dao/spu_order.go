package dao

import (
	"flash-sale-backend/model"
)

func CreateOrder(m *model.SPUOrder) error {
	if err := conn.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOrderState(id int64, state int32) {
	conn.Table("spu_order").Where("id=?", id).Update("state=?", state)
}
