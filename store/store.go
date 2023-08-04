package store

import (
	helper "inmem-db/helpers"
	"sync"
)

const defaultSize = 50

type HashTable struct {
	maps []sync.Map
	size int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		maps: make([]sync.Map, size),
		size: size,
	}
}

func (hashtable *HashTable) Add(key string, value interface{}) {
	hash := helper.Hash(key)
	hashtable.maps[hash].Store(key, value)
}

func (hashtable *HashTable) Get(key string) (any, bool) {
	hash := helper.Hash(key)
	value, ok := hashtable.maps[hash].Load(key)
	if !ok {
		return "", false
	}
	return value, true
}

func (hashtable *HashTable) Del(key string) {
	hash := helper.Hash(key)
	hashtable.maps[hash].Delete(key)
}
