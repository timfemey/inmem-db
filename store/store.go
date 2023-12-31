package store

import (
	"bytes"
	"fmt"
	helper "inmem-db/helpers"
	"sync"
)

type HashTable struct {
	maps     []sync.Map
	size     int64
	compress bool
}

func NewHashTable(size int64, compress bool) *HashTable {
	return &HashTable{
		maps:     make([]sync.Map, size),
		size:     size,
		compress: compress,
	}
}

func (hashtable *HashTable) Add(key string, value interface{}) error {
	hash := helper.Hash(key) % uint64(hashtable.size)
	if hashtable.compress {
		compressedData, err := helper.Compress(value)
		if err != nil {
			return err
		}
		hashtable.maps[hash].Store(key, compressedData)

	} else {
		hashtable.maps[hash].Store(key, value)
	}
	return nil

}

func (hashtable *HashTable) Get(key string) (any, bool) {
	hash := (helper.Hash(key)) % uint64(hashtable.size)
	value, ok := hashtable.maps[hash].Load(key)
	if !ok {
		return "", false
	}
	if hashtable.compress {
		val := value.(bytes.Buffer)
		unCompressedData, err := helper.DeCompress(&val)
		if err != nil {
			fmt.Printf("%+vn", err)
			return "", false
		}
		return unCompressedData, true
	}
	return value, true

}

func (hashtable *HashTable) Del(key string) {
	hash := helper.Hash(key)
	hashtable.maps[hash].Delete(key)
}
