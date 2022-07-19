package dao

import "flash-sale-backend/model"

func CreateSPUItem(m *model.SPUItem) {
	conn.Create(m)
}

func GetAllSPUInfo() []*model.SPUItem {
	data := make([]*model.SPUItem, 0)
	if err := conn.Find(&data).Error; err != nil {
		return data
	}
	return data
}

func GetSpuItemByID(id []int64) map[int64]*model.SPUItem {
	itemList := make([]*model.SPUItem, 0)
	m := make(map[int64]*model.SPUItem)
	if err := conn.Where("id in ?", id).Find(&itemList).Error; err != nil {
		return m
	}

	for _, item := range itemList {
		m[item.ID] = item
	}
	return m
}
