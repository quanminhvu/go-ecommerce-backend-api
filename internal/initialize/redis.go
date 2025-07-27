package initialize

import (
	"context"
	"fmt"

	"github.com/quanminhvu/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // pool size
	})

	_, err := rdb.Ping(ctx).Result()
	checkErr(err, "InitRedis error")
	fmt.Println("Redis connected")
	global.Logger.Info("Redis connected", zap.String("ok", "success"))
	global.Rdb = rdb
}
