package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nsrvel/go-fiber-boilerplate/config"
)

var RDB *redis.Client

func NewRedisClient(conf *config.RedisAccount) *redis.Client {
	redisHost := conf.Host

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Password:     conf.Password,
		MinIdleConns: conf.MinIdleConns,
		PoolSize:     conf.PoolSize,
		PoolTimeout:  conf.PoolTimeout,
		DB:           conf.DB,
	})

	RDB = client
	return client
}

func RedisSet(key string, data interface{}, exp time.Duration) error {

	//* store data using SET command
	op1 := RDB.Set(context.Background(), key, data, exp)
	if err := op1.Err(); err != nil {
		return errors.New(fmt.Sprintf("Unable to SET data. error: %v", err))
	}
	return nil
}

func RedisGet(key string) (error, string) {

	//* get data using Get command
	op2 := RDB.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		return errors.New(fmt.Sprintf("Unable to GET data. error: %v", err)), ""
	}
	res, err := op2.Result()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to GET data. error: %v", err)), ""
	}
	return nil, res
}

func RedisDel(key string) error {

	//* delete key using Del command
	op2 := RDB.Del(context.Background(), key)
	if err := op2.Err(); err != nil {
		return errors.New(fmt.Sprintf("Unable to DELETE data. error: %v", err))
	}
	return nil
}
