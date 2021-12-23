package main

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	//l := list.New()
	//l.PushBack(4)
	//l.PushBack(1)
	//e := l.Front()
	//fmt.Println(e.Value)

	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	res := cache.Get(1)
	fmt.Println(res)
	res = cache.Get(2)
	fmt.Println(res)
}
