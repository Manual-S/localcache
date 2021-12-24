package main

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	capacity := 2
	lru := NewLRUCache(capacity)
	lru.Put("1", Item{
		Key:    "1",
		Object: "1",
	})

	lru.Put("2", Item{
		Key:    "2",
		Object: "2",
	})

	lru.Put("3", Item{
		Key:    "3",
		Object: "3",
	})

	value := lru.Get("1")
	fmt.Println(value)
}
