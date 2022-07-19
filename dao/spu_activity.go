package dao

import (
	"context"
	"flash-sale-backend/model"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func CreateSPUActivity(m *model.SPUActivity, total int64) {
	conn.Create(m)
	redisConn.Set(context.Background(), fmt.Sprintf("activity:end:%v", m.ID), m.EndTime.Unix(), time.Duration(total)*time.Second)
	redisConn.Set(context.Background(), fmt.Sprintf("activity:stock:%v", m.ID), m.AvailStock, time.Duration(total)*time.Second)
}

func SetStock(id int64, stock int32) {
	redisConn.Set(context.Background(), fmt.Sprintf("activity:stock:%v", id), stock, 10*time.Minute)
}

func GetAllSPUActivity() []*model.SPUActivity {
	data := make([]*model.SPUActivity, 0)
	now := time.Now()
	if err := conn.Where("start_time<?", now).Where("end_time>?", now).Find(&data).Error; err != nil {
		return data
	}

	return data
}

func GetActivityByID(id int64) *model.SPUActivity {
	data := new(model.SPUActivity)
	if err := conn.Where("id=?", id).Take(data).Error; err != nil {
		return nil
	}

	return data
}

func GetActivityEndTimeByID(id int64) int64 {
	key := fmt.Sprintf("activity:end:%v", id)
	result := redisConn.Get(context.Background(), key)
	if result.Err() != nil {
		m := GetActivityByID(id)
		if m != nil {
			redisConn.Set(context.Background(), fmt.Sprintf("activity:end:%v", m.ID), m.EndTime.Unix(), m.EndTime.Sub(m.StartTime))
			redisConn.Set(context.Background(), fmt.Sprintf("activity:stock:%v", m.ID), m.AvailStock, m.EndTime.Sub(m.StartTime))
		}

	}
	endTime, _ := result.Int64()

	return endTime

}

func LockStockByID(id int64) int64 {
	updates := map[string]interface{}{
		"avail_stock": gorm.Expr("avail_stock-1"),
		"lock_stock":  gorm.Expr("lock_stock+1"),
	}
	result := conn.Table("spu_activity").Where("id = ?", id).Where("avail_stock>?", 0).Updates(updates)
	if result.Error != nil {
		return -1
	}
	rows := result.RowsAffected
	return rows
}

func CheckAvalStock(id int64) bool {
	scriptStr := "if redis.call('exists',KEYS[1]) == 1 then\n" +
		"                 local stock = tonumber(redis.call('get', KEYS[1]))\n" +
		"                 if( stock <=0 ) then\n" +
		"                    return -1\n" +
		"                 end;\n" +
		"                 redis.call('decr',KEYS[1]);\n" +
		"                 return stock - 1;\n" +
		"             end;\n" +
		"             return -1;"
	script := redis.NewScript(scriptStr)
	key := fmt.Sprintf("activity:stock:%v", id)
	result := script.Run(context.Background(), redisConn, []string{key})
	if result.Err() != nil {
		return false
	}

	if i, _ := result.Int64(); i >= 0 {
		return true
	}

	return false
}
