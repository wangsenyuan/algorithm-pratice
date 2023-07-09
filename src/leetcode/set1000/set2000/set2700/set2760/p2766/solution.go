package p2766

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	pos := make(map[int]int)

	for _, v := range nums {
		pos[v]++
	}

	for i := 0; i < len(moveFrom); i++ {
		x := moveFrom[i]
		y := moveTo[i]
		delete(pos, x)
		pos[y]++
	}

	ans := make([]int, 0, len(pos))

	for k := range pos {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	return ans
}
