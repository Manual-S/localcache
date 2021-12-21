package main

import (
	"sync"
	"time"
)

type Item struct {
	Object     interface{}
	Expiration int64 // 过期时间
}

// Expired 过期返回true 不过期返回false
func (item Item) Expired() bool {
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
}

type cache struct {
	data map[string]Item
	mu   sync.RWMutex // go提供的读写锁
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

func (c *Cache) Set(key string, value Item) {
	c.mu.Lock()
	c.data[key] = Item{
		Object:     value.Object,
		Expiration: value.Expiration,
	}
	c.mu.Unlock()
}

func NewCache() *Cache {
	items := make(map[string]Item)
	c := &cache{
		data: items,
	}
	return &Cache{c}
}
