package model

import "time"

type SPUActivity struct {
	ID         int64
	ActName    string
	SpuId      int64
	SalePrice  float32
	ActStatus  int8
	Prime      int8
	Special    int8
	StartTime  time.Time
	EndTime    time.Time
	TotalStock int32
	AvailStock int32
	LockStock  int32
}

func (SPUActivity) TableName() string {
	return "spu_activity"
}
