package timecache

import "time"

var (
	//DefaultLiveTime 默认存活时间
	DefaultLiveTime time.Duration = time.Hour * 24
)

//FOREVER 永不过期
const FOREVER = time.Hour * 24 * 365 * 100
