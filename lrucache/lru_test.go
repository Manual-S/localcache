package lrucache

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkLRU_Rand lrucache的命中率
func BenchmarkLRU_Rand(b *testing.B) {
	capacity := 8192
	l, err := NewLRUCache(capacity)
	assert.Nil(b, err)
	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = rand.Int63() % 32768
	}
	b.ResetTimer() // 重置计数器 可以用来忽略一些准备工作
	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		key := strconv.FormatInt(trace[i], 10)
		if i%2 == 0 {
			l.Put(key, Item{
				Key:    key,
				Object: key}, 0)
		} else {
			_, ok := l.Get(key)
			if ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}
