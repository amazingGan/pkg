package xgo

import (
	"sync"
	"sync/atomic"
	"time"
)

// leap window
type Carrier struct {
	windows  map[string]*window
	m        sync.Mutex
	duration time.Duration
}

func NewCarrier(winDuration time.Duration) *Carrier {
	return &Carrier{
		duration: winDuration,
		windows:  make(map[string]*window),
	}
}

// GetWindow 返回window，如果window不存在, 创建一个window
// 如果window存在，则根据window的dead(end) timeUtil，滚动窗口
// 潜在风险: key没有过期，某些场景c.windows会非常大
// todo(lvchao):
// 1. 自动清除过期(冷)key
// 2. 自动识别热点，避免偶发的key触发大量的window计算 (LRU?)
// 3. 引入对象池, 减少gc压力
func (c *Carrier) GetWindow(key string) *window {
	c.m.Lock()
	defer c.m.Unlock()

	if win, ok := c.windows[key]; ok {
		if atomic.LoadInt64(&win.dead) <= time.Now().Unix() {
			win.Roll()
		}

		return win
	}

	win := &window{
		duration: c.duration,
		dead:     time.Now().Add(c.duration).Unix(),
	}
	c.windows[key] = win
	return win
}

type window struct {
	done uint32

	duration time.Duration
	dead     int64
}

// 在一个窗口期内，保证成功执行一次
func (win *window) DoOnce(f func() bool) {
	if atomic.LoadInt64(&win.dead) > time.Now().Unix() {
		if atomic.CompareAndSwapUint32(&win.done, 0, 1) {
			if f() {
				atomic.StoreInt64(&win.dead, time.Now().Add(win.duration).Unix())
			} else {
				atomic.StoreUint32(&win.done, 0)
			}
		}
	}
}

// Roll 滚动窗口
func (win *window) Roll() {
	atomic.StoreInt64(&win.dead, time.Now().Add(win.duration).Unix())
	atomic.StoreUint32(&win.done, 0)
}
