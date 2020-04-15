// Copyright 2020. All rights reserved.
// Author 赵路通

package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"time"
)

var (
	redisAddr string
	redisPwd  string
)

// 返回一个redis客户端
func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPwd,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     20,
		DB:           2,
	})
	return client
}

// 使用键值设置
func SetValueKeyWithExpire(key string, data interface{}, expired time.Duration) error {
	client := redisClient()
	defer client.Close()
	s := client.Set(key, data, expired)
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

// 批量设置hash值，无过期时间
func MultHashSetWithoutExpire(key string, values map[string]interface{}) error {
	client := redisClient()
	defer client.Close()
	s := client.HMSet(key, values)
	if err := s.Err(); err != nil {
		logs.Error("批量设置哈希值错误", err)
		return err
	}
	return nil
}

// 从redis里面读取hash值
func HgetValue(key string, field string) (string, error) {
	client := redisClient()
	defer client.Close()
	s := client.HGet(key, field)
	r, err := s.Result()
	if err != nil {
		logs.Error("获取哈希数据错误", err)
		return "", err
	}
	return r, nil
}

// 从redis中根据键获取值
func GetValue(key string) string {
	client := redisClient()
	defer client.Close()
	v, _ := client.Get(key).Result()
	return v
}

// 新增值类型
func IncrKey(key string) int64 {
	client := redisClient()
	defer client.Close()
	v, _ := client.Incr(key).Result()
	return v
}

// 设置redis的过期时间
func ExpiredKey(key string, time time.Duration) (bool, error) {
	client := redisClient()
	defer client.Close()
	v, err := client.Expire(key, time).Result()
	return v, err
}

// 查询是否有某个值
func KeyExists(key string) int64 {
	client := redisClient()
	defer client.Close()
	v, _ := client.Exists(key).Result()
	return v
}

// 新增Hash中的某个值
func HKeyIncrBy(key string, field string, incr int64) int64 {
	client := redisClient()
	defer client.Close()
	v, _ := client.HIncrBy(key, field, incr).Result()
	return v
}

// 从redis中删除数据
func DelValue(key string) (int64, error) {
	client := redisClient()
	defer client.Close()
	ssc := client.Del(key)
	return ssc.Result()
}

func Keys(patterm string) ([]string, error) {
	client := redisClient()
	defer client.Close()
	sc := client.Keys(patterm)
	if err := sc.Err(); err != nil {
		logs.Error("模糊查询键值错误", err)
		return nil, err
	}
	return sc.Result()
}
