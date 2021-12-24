package cache

import (
	"cache/lrucache"
	"fmt"
	"testing"
	"time"
)

func TestCache_Set(t *testing.T) {
	cache := NewCache()
	for i := 0; i < 10; i++ {
		go func(i int) {
			key := "key"
			cache.Set(key, lrucache.Item{
				Object: i,
			}, 0)
		}(i)
	}

	time.Sleep(5 * time.Second)

	key := "key"
	value, _ := cache.Get(key)
	fmt.Println(value)

}
