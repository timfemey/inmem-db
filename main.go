package main

import (
	"flag"
	"inmem-db/store"
)

var initialSize int64
var compress bool
var hashtable store.HashTable

func init() {
	sizeArg := flag.Int("size", 1000, "How many Entries Memory DB should allow, Default is 1000")
	compressArg := flag.Bool("compress", false, "Compress Values sent")

	flag.Parse()

	initialSize = int64(*sizeArg)
	compress = *compressArg

	hashtable = *store.NewHashTable(initialSize, compress)

}

func main() {

}
