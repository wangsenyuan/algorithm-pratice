package main

import (
	"container/heap"
	"fmt"
)

func main() {
	cache := Constructor(2)
	cache.Set(1, 1)
	cache.Set(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Set(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3.
	cache.Set(4, 4)           // evicts key 1.
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}

type LFUCache struct {
	pq      *PriorityQueue
	itemMap map[int]*Item
	index   int
}

func Constructor(capacity int) LFUCache {
	pq := make(PriorityQueue, 0, capacity)
	heap.Init(&pq)
	return LFUCache{&pq, make(map[int]*Item), 0}
}

func (this *LFUCache) Get(key int) int {
	if item, ok := this.itemMap[key]; ok {
		this.index++
		item.pos = this.index
		item.freq++
		this.pq.update(item)
		return item.value
	}
	return -1
}

func (this *LFUCache) Set(key int, value int) {
	if cap(*this.pq) == 0 {
		return
	}
	item, ok := this.itemMap[key]
	if !ok && len(*this.pq) == cap(*this.pq) {
		item = heap.Pop(this.pq).(*Item)
		//fmt.Printf("item {%d, %d, %d, %d} is poped out\n", item.key, item.value, item.freq, item.pos)
		delete(this.itemMap, item.key)
	}

	if ok {
		this.index++
		item.pos = this.index
		item.freq++
		item.value = value
		this.pq.update(item)
	} else {
		this.index++
		item = &Item{key, value, 1, this.index, 0}
		this.itemMap[key] = item
		heap.Push(this.pq, item)
	}
}

type Item struct {
	key, value, freq, pos, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq < pq[j].freq || (pq[i].freq == pq[j].freq && pq[i].pos < pq[j].pos) {
		return true
	}
	return false
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}
