package p1424

import "sort"

func findDiagonalOrder(nums [][]int) []int {
	items := make([]Item, 0, 10000)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			items = append(items, Item{nums[i][j], i, j})
		}
	}

	sort.Sort(Items(items))

	res := make([]int, len(items))

	for i := 0; i < len(items); i++ {
		res[i] = items[i].val
	}

	return res
}

type Item struct {
	val int
	x   int
	y   int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	if items[i].x+items[i].y < items[j].x+items[j].y {
		return true
	}

	if items[i].x+items[i].y == items[j].x+items[j].y {
		return items[i].x > items[j].x
	}

	return false
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
