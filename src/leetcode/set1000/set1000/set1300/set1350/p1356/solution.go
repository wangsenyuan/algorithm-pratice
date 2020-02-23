package p1356

import (
	"sort"
)

func sortByBits(arr []int) []int {
	n := len(arr)
	items := make([]Item, n)

	for i := 0; i < n; i++ {
		num := arr[i]
		var cnt int
		for num > 0 {
			x := num & 1
			num >>= 1
			cnt += x
		}

		items[i] = Item{arr[i], cnt}
	}

	sort.Sort(Items(items))

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = items[i].num
	}

	return res
}

type Item struct {
	num int
	cnt int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].cnt < items[j].cnt || items[i].cnt == items[j].cnt && items[i].num < items[j].num
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
