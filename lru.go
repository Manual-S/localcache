// 实现lru算法 参考leetcode146
package main

import (
	"container/list"
	"sync"
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

func NewLRUCache(capacity int) LRUCache {
	hash := make(map[string]*list.Element)
	return LRUCache{
		capacity: capacity,
		hash:     hash,
		cache:    list.New(),
	}
}

// Get 查看map中是否存在该元素
func (c *LRUCache) Get(key string) (interface{}, bool) {
	it, ok := c.hash[key]

	if !ok {
		return nil, false
	}

	// 操作list
	c.cache.MoveBefore(it, c.cache.Front())
	return it.Value.(Item).Object, true
}

func (c *LRUCache) Put(key string, value Item) {
	it, ok := c.hash[key]

	if ok {
		// key已经存在  更新key值
		it.Value = Item{
			Key:        value.Key,
			Object:     value.Object,
			Expiration: value.Expiration,
		}
		// 调整list
		// c.cache.MoveAfter(it, c.cache.Front())
		c.cache.MoveToFront(it)
		// 返回
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
}
