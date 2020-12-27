package p1704

import "container/heap"

const MAX_DAY = 50000

func eatenApples(apples []int, days []int) int {
	n := len(apples)

	que := make(Items, 0, n)

	var res int

	for i := 0; i < MAX_DAY; i++ {
		for que.Len() > 0 && que[0].rotten <= i {
			heap.Pop(&que)
		}

		if i < n && apples[i] > 0 {
			item := new(Item)
			item.count = apples[i]
			item.rotten = i + days[i]
			heap.Push(&que, item)
		}

		if i > n && que.Len() == 0 {
			break
		}

		if que.Len() > 0 {
			cur := heap.Pop(&que).(*Item)
			cur.count--
			if cur.count > 0 {
				heap.Push(&que, cur)
			}
			res++
		}
	}

	return res
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */

type Item struct {
	count  int
	rotten int
	index  int
}

type Items []*Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].rotten < items[j].rotten
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
