package p347

import "sort"

type tuple2 struct {
	val int
	cnt int
}

type tuple2List []*tuple2

func (list tuple2List) Len() int {
	return len(list)
}

func (list tuple2List) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	return a.cnt <= b.cnt
}

func (list tuple2List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)

	list := make([]*tuple2, 0, len(nums))

	prev := &tuple2{nums[0], 1}

	list = append(list, prev)

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x == prev.val {
			prev.cnt++
			continue
		}

		prev = &tuple2{x, 1}
		list = append(list, prev)
	}

	sort.Sort(sort.Reverse(tuple2List(list)))

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		t := list[i]
		result = append(result, t.val)
	}

	return result
}
