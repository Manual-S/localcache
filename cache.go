package main

import (
	"sync"
	"time"
)

type Item struct {
	Object     interface{}
	Expiration int64 // 过期时间 如果过期时间为0 则不过期
}

// Expired 过期返回true 不过期返回false
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
}

type cache struct {
	data map[string]Item // 存放缓存的底层结构
	mu   sync.RWMutex    // go提供的读写锁

}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.data[key]
	if !found {
		// 没有找到数据
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()
	return item.Object, true
}

func (c *Cache) Set(key string, value Item, d time.Duration) {
	var end int64
	if d > 0 {
		end = time.Now().Add(d).UnixNano()
	}

	c.mu.Lock()
	c.data[key] = Item{
		Object:     value.Object,
		Expiration: end,
	}
	c.mu.Unlock()
}

func (c *Cache) set(key string, value interface{}, d time.Duration) {
	var end int64
	if d > 0 {
		end = time.Now().Add(d).UnixNano()
	}
	c.data[key] = Item{
		Object:     value,
		Expiration: end,
	}
}

func NewCache() *Cache {
	items := make(map[string]Item)
	c := &cache{
		data: items,
	}
	return &Cache{c}
}
