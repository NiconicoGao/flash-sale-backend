package model

import "time"

type SPUOrder struct {
	ID         int64
	OrderId    int64
	ActivityId int64
	State      int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (SPUOrder) TableName() string {
	return "spu_order"
}
