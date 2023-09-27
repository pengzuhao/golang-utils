package redisopera

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisClient(ipaddr, passwd string) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     ipaddr,
		Password: passwd,
		DB:       0, // use default DB
	})
	return rdb, err
}

func RedisSet(ipaddr, passwd, key string, value interface{}, expired time.Duration) (err error) {
	ctx := context.Background()
	rdb, err := RedisClient(ipaddr, passwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rdb.Set(ctx, key, value, expired).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisGet(ipaddr, passwd, key string) (val interface{}, err error) {
	ctx := context.Background()
	rdb, err := RedisClient(ipaddr, passwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err = rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		return
	}
}
func RedisConnPoolGet(ipaddr, passwd, key string) (val interface{}, err error) {
	ctx := context.Background()
	rdb, err := RedisClient(ipaddr, passwd)
	if err != nil {
		return
	}
	conn := rdb.Conn()
	defer conn.Close()
	val, err = conn.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		return
	}
}
