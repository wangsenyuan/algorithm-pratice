package p0003

import "container/heap"

type FirstUnique struct {
	cur int
	pq  *Items
	mem map[int]*Item
}

func Constructor(nums []int) FirstUnique {
	mem := make(map[int]*Item)

	for i := 0; i < len(nums); i++ {
		if v, found := mem[nums[i]]; !found {
			mem[nums[i]] = &Item{nums[i], 1, i, i}
		} else {
			v.count++
		}
	}

	items := make(Items, len(mem))

	var cur int

	for _, v := range mem {
		// v.occureAt = cur
		v.index = cur
		items[cur] = v
		cur++
	}

	heap.Init(&items)

	return FirstUnique{len(nums), &items, mem}
}

func (this *FirstUnique) ShowFirstUnique() int {
	pq := *(this.pq)

	if len(pq) == 0 {
		return -1
	}
	res := pq[0]

	if res.count == 1 {
		return res.value
	}

	return -1
}

func (this *FirstUnique) Add(value int) {
	if v, found := this.mem[value]; found {
		this.pq.addCount(v)
	} else {
		item := new(Item)
		item.value = value
		item.occureAt = this.cur
		item.count = 1
		heap.Push(this.pq, item)
		this.mem[value] = item
	}

	this.cur++
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */

type Item struct {
	value    int
	count    int
	occureAt int
	index    int
}

type Items []*Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].count < items[j].count || items[i].count == items[j].count && items[i].occureAt < items[j].occureAt
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
	items[i].index = i
	items[j].index = j
}

func (items *Items) Push(ele interface{}) {
	n := len(*items)

	item := ele.(*Item)
	item.index = n

	*items = append(*items, item)
}

func (items *Items) Pop() interface{} {
	old := *items
	n := len(old)
	res := old[n-1]
	res.index = -1
	*items = old[:n-1]
	return res
}

func (items *Items) addCount(item *Item) {
	item.count++
	heap.Fix(items, item.index)
}
