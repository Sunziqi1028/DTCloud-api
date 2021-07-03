/**
* @Author: lik
* @Date: 2021/3/8 14:15
* @Version 1.0
 */
package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheAdapter *cache.Cache

func init() {
	// 默认过期时间为5min，每10min清理一次过期缓存
	cacheAdapter = cache.New(5*time.Minute, 10*time.Second)
}

func SetCache(k string, x interface{}, d time.Duration) {
	cacheAdapter.Set(k, x, d)
}

func GetCache(k string) (interface{}, bool) {
	return cacheAdapter.Get(k)
}

//设置cache 无时间参数
func SetDefaultCache(k string, x interface{}) {
	cacheAdapter.SetDefault(k, x)
}

//删除 cache
func DeleteCache(k string) {
	cacheAdapter.Delete(k)
}

// Add() 加入缓存
func AddCache(k string, x interface{}, d time.Duration) {
	cacheAdapter.Add(k, x, d)
}

// IncrementInt() 对已存在的key 值自增n
func IncrementIntCache(k string, n int) (num int, err error) {
	return cacheAdapter.IncrementInt(k, n)
}
