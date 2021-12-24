// 实现lru算法 参考leetcode146
package main

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

// Item 存储实体
type Item struct {
	Key        string
	Object     interface{}
	Expiration int64 // 过期时间 如果过期时间为0 则不过期
}

type LRUCache struct {
	capacity int
	hash     map[string]*list.Element
	cache    *list.List
	lock     sync.RWMutex
}

func NewLRUCache(capacity int) (LRUCache, error) {
	if capacity <= 0 {
		return LRUCache{}, errors.New("capacity size < 0 ")
	}
	hash := make(map[string]*list.Element)
	return LRUCache{
		capacity: capacity,
		hash:     hash,
		cache:    list.New(),
	}, nil
}

// Get 查看map中是否存在该元素
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	it, ok := c.hash[key]

	if !ok {
		return nil, false
	}

	// 操作list
	c.cache.MoveBefore(it, c.cache.Front())
	c.lock.RUnlock()
	return it.Value.(Item).Object, true
}

func (c *LRUCache) Put(key string, value Item, d time.Duration) {
	c.lock.Lock()

	it, ok := c.hash[key]

	var end int64
	if d > 0 {
		end = time.Now().Add(d).UnixNano()
	}

	if ok {
		// key已经存在  更新key值
		it.Value = Item{
			Key:        value.Key,
			Object:     value.Object,
			Expiration: end,
		}
		// 调整list
		// c.cache.MoveAfter(it, c.cache.Front())
		c.cache.MoveToFront(it)
		// 返回
		c.lock.Unlock()
		return
	}

	if c.cache.Len() == c.capacity {
		// 容量到达上限 list需要做删除
		temp := c.cache.Back()
		c.cache.Remove(temp)
		// hash也要做删除
		delete(c.hash, temp.Value.(Item).Key)
	}

	// 正常添加元素
	c.cache.PushFront(Item{
		Key:        value.Key,
		Object:     value.Object,
		Expiration: value.Expiration,
	})

	c.hash[key] = c.cache.Front()
	c.lock.Unlock()
}

func (c *LRUCache) Len() int {
	c.lock.RLock()
	length := c.cache.Len()
	c.lock.RUnlock()
	return length
}
