package timecache

import "time"

// Add 添加一个缓存
// v: 缓存关联的值
// timestamp: 有效期
func Add(v interface{}, timestamp time.Duration) (key string, err error) {
	if cache == nil {
		return "", ErrCacheNotInit
	}
	return cache.Add(v, timestamp)
}

// Set 更新缓存
func Set(key string, value interface{}, timestamp time.Duration) error {
	if cache == nil {
		return ErrCacheNotInit
	}
	return cache.Set(key, value, timestamp)
}

// Get 获取缓存
// key: 缓存key值
// fresh：是否刷新有效期(重新计时)
// bool: 是否存在
// interface{}: 关联值
// error: 错误信息
func Get(key string, v interface{}, timestamp time.Duration) (bool, error) {
	if cache == nil {
		return false, ErrCacheNotInit
	}
	return cache.Get(key, v, timestamp)
}

// Delete 删除缓存
func Delete(key string) error {
	if cache == nil {
		return ErrCacheNotInit
	}
	return cache.Delete(key)
}

var cache Cache

// Cache 缓存接口
type Cache interface {
	Add(interface{}, time.Duration) (string, error)
	Set(string, interface{}, time.Duration) error
	Get(string, interface{}, time.Duration) (bool, error)
	Delete(string) error
	Eval(string, []string, ...interface{}) (interface{}, error)
}

// InitInfo 初始化数据结构
type InitInfo struct {
	Tp            string        //redis
	Addr          string        //地址
	Password      string        //密码
	DB            int           //数据库
	ReaderTimeout time.Duration //读取超时
	WriteTimeout  time.Duration //写入超时
	PoolSize      int           //连接池数量
}

// Init 初始化函数
func Init(ii *InitInfo) (Cache, error) {
	if cache != nil {
		return cache, nil
	}
	var err error
	if cache == nil {
		cache, err = initRedis(ii)
	}
	return cache, err
}

// NewCache 生成一个新的redis cli
func NewCache(ii *InitInfo) (Cache, error) {
	return initRedis(ii)
}

func DefaultCache() Cache {
	if cache == nil {
		panic("Cache 未初始化")
	}
	return cache
}
