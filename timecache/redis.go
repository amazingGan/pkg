package timecache

import (
	"encoding"
	"encoding/json"
	"time"

	"github.com/amazingGan/pkg/uuid"
	"github.com/go-redis/redis"
)

func initRedis(ii *InitInfo) (Cache, error) {
	client := &redisCache{}
	client.info = ii
	client.client = redis.NewClient(&redis.Options{
		Addr:         ii.Addr,
		Password:     ii.Password, // no password set
		DB:           ii.DB,       // use default DB
		ReadTimeout:  ii.ReaderTimeout,
		WriteTimeout: ii.WriteTimeout,
		PoolSize:     ii.PoolSize,
	})
	err := client.client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return client, err
}

type redisCache struct {
	client *redis.Client
	info   *InitInfo
}

func (cache *redisCache) Set(key string, value interface{}, timestamp time.Duration) error {
	if timestamp <= 0 {
		timestamp = DefaultLiveTime
	}
	return cache.client.Set(key, cache.marshal(value), timestamp).Err()
}

func (cache *redisCache) Add(value interface{}, timestamp time.Duration) (string, error) {
	key := uuid.UUID32() //防止共用服务导致的key冲突
	return key, cache.Set(key, value, timestamp)
}

func (cache *redisCache) marshal(v interface{}) interface{} {
	switch v.(type) {
	case nil, string, []byte, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, encoding.BinaryMarshaler:
		return v
	default:
		return &redisValue{v: v}
	}
}

type redisValue struct {
	v interface{}
}

func (rv *redisValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(rv.v)
}

func (rv *redisValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, rv.v)
}

// Get 获取值
// key: 存储key
// value: 存储value
// timestamp: 是否更新有效期
// bool: 标记是否存在
// error: 错误信息
func (cache *redisCache) Get(key string, value interface{}, timestamp time.Duration) (bool, error) {
	result := cache.client.Get(key)
	if result.Err() != nil {
		if result.Err() == redis.Nil {
			return false, nil
		}
		return false, result.Err()
	}
	var err error
	switch value.(type) {
	case nil, *string, *[]byte, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64, *bool, encoding.BinaryMarshaler:
		err = result.Scan(value)
	default:
		v := &redisValue{v: value}
		err = result.Scan(v)
	}
	if err != nil {
		return true, err
	}
	if timestamp > 0 { //重置有效期
		result := cache.client.Expire(key, timestamp)
		if result.Err() != nil {
			return true, err
		}
	}
	return true, nil
}

func (cache *redisCache) Delete(key string) error {
	result := cache.client.Del(key)
	if result.Err() == redis.Nil {
		return nil
	}
	return result.Err()
}

func (cache *redisCache) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return cache.client.Eval(script, keys, args...).Result()
}

func (cache *redisCache) GetClient() *redis.Client {
	return cache.client
}
