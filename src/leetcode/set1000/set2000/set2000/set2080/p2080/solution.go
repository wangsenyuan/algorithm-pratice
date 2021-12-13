package p2080

import "sort"

type RangeFreqQuery struct {
	nums map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	nums := make(map[int][]int)
	update := func(v, p int) {
		if _, ok := nums[v]; !ok {
			nums[v] = make([]int, 0, 2)
		}
		nums[v] = append(nums[v], p)
	}

	for i, v := range arr {
		update(v, i)
	}

	return RangeFreqQuery{nums}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if vs, ok := this.nums[value]; !ok {
		return 0
	} else {
		i := sort.Search(len(vs), func(i int) bool {
			return vs[i] > right
		})
		j := sort.Search(len(vs), func(j int) bool {
			return vs[j] >= left
		})
		return i - j
	}
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
