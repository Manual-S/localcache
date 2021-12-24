package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	capacity := 2
	lru, err := NewLRUCache(capacity)
	assert.Nil(t, err)

}
