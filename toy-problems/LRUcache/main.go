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
	value := -1
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

	if this.isFull {
		i := this.size - 1
		this.storage = append(this.storage, item)
		this.storage = append(this.storage[:i], this.storage[i+1:]...)
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
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println("Cache is", cache.Test())

	cache.Get(1)
	cache.Put(3, 3)
	fmt.Println("Cache is", cache.Test())

	cache.Get(2)
	cache.Put(4, 4)
	fmt.Println("Cache is", cache.Test())

	cache.Get(3)
	cache.Get(4)
	fmt.Println("Cache is", cache.Test())

}
