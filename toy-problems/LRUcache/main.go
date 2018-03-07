package main

import "fmt"

type Item struct {
	key   int
	value int
}

type LRUCache struct {
	storage []Item
	isFull  bool
	size    int
	limit   int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		storage: []Item{},
		isFull:  false,
		size:    0,
		limit:   capacity,
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	var value int
	for i, item := range this.storage {
		if item.key == key {
			if i != 0 {
				this.storage[i-1], this.storage[i] = this.storage[i], this.storage[i-1]
			}
			value = item.value
			break
		}
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {
	item := Item{key, value}

	if this.isFull == true {
		i := this.size - 1
		this.storage[i] = item
	} else {
		this.size++
		this.storage = append(this.storage, item)

		if this.size == this.limit {
			this.isFull = true
		}
	}
}

func (this *LRUCache) Test() []Item {
	return this.storage
}

func main() {
	cache := Constructor(2)

	cache.Put(2, 10)
	cache.Put(1, 5)

	cache.Get(1)

	cache.Put(3, 5)

	fmt.Printf("Get 2 should return 10: %v\n", cache.Get(2))
	fmt.Printf("Show the storage: %v\n", cache.Test())
}
