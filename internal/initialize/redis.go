package initialize

import (
	"context"
	"ecommerce/global"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis(){
	r:= global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
        Password: r.Password, // no password set
        DB:       r.Database,  // use default DB
		PoolSize: 10,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	global.Redis = rdb
}


