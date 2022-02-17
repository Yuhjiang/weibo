package database

import (
	"github.com/Yuhjiang/weibo/global"
	"github.com/go-redis/redis/v8"
	"time"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

type RedisStore struct {
	Prefix       string
	DefaultValue interface{}
}

func (rs *RedisStore) Set(key, value string, exp time.Duration) error {
	err := Redis.Set(global.Ctx, rs.Prefix+key, value, exp).Err()
	return err
}

func (rs *RedisStore) Get(key string) string {
	res, err := Redis.Get(global.Ctx, rs.Prefix+key).Result()
	if err != nil {
		return rs.DefaultValue.(string)
	} else {
		return res
	}
}
