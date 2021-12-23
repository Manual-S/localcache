// 实现lru算法 参考leetcode146
package main

import (
	"container/list"
)

type Node struct {
	Key   int
	Value int
}

type LRUCache struct {
	capacity int
	m        map[int]*list.Element
	cache    list.List
}

func Constructor(capacity int) LRUCache {
	hash := make(map[int]*list.Element)
	return LRUCache{
		capacity: capacity,
		m:        hash,
	}
}

// Get 查看map中是否存在该元素
func (c *LRUCache) Get(key int) int {
	it, ok := c.m[key]
	if !ok {
		return -1
	}

	// 操作list
	c.cache.MoveBefore(it, c.cache.Front())
	return it.Value.(Node).Value
}

func (c *LRUCache) Put(key int, value int) {
	it, ok := c.m[key]

	if ok {
		// key已经存在  更新key值
		it.Value = Node{
			Key:   key,
			Value: value,
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
		delete(c.m, temp.Value.(Node).Key)
	}

	// 正常添加元素
	c.cache.PushFront(Node{
		Key:   key,
		Value: value,
	})

	c.m[key] = c.cache.Front()
}
