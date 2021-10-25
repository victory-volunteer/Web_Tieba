package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
	//Nil = redis.Nil
)

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}
func Close() {
	_ = rdb.Close() //因为上面rdb变量不想让外界访问，所以在此处定义一个方法用来释放rdb对象
}
