package model

import "time"

type SPUItem struct {
	ID        int64
	SPUName   string  `gorm:"column:spu_name"`
	SPUPrice  float32 `gorm:"column:spu_price"`
	SPUDesc   string  `gorm:"column:spu_desc"`
	SPUImage  string  `gorm:"column:spu_image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SPUItem) TableName() string {
	return "spu_item"
}
