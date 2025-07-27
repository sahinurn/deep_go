package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.True(t, data.Contains(5))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))
	assert.False(t, data.Contains(15))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestOrderedMapExtra(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Zero(t, data.Size())

	data.Insert(33, 33)
	data.Insert(5, 5)
	data.Insert(35, 35)
	data.Insert(99, 99)
	data.Insert(1, 1)
	data.Insert(4, 4)
	data.Insert(20, 20)
	data.Insert(17, 17)
	data.Insert(31, 31)

	assert.Equal(t, 9, data.Size())
	assert.True(t, data.Contains(33))
	assert.True(t, data.Contains(99))
	assert.True(t, data.Contains(17))
	assert.True(t, data.Contains(20))
	assert.True(t, data.Contains(5))
	assert.True(t, data.Contains(31))

	data.Erase(5)

	assert.Equal(t, 8, data.Size())
	assert.True(t, data.Contains(17))
	assert.True(t, data.Contains(20))
	assert.True(t, data.Contains(31))
	assert.False(t, data.Contains(5))
}
