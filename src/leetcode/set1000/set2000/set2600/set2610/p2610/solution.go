package p2610

func findMatrix(nums []int) [][]int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	var r int
	for _, v := range cnt {
		r = max(r, v)
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, 0, len(cnt))
		ct := make(map[int]int)
		for k, v := range cnt {
			res[i] = append(res[i], k)
			if v > 1 {
				ct[k] = v - 1
			}
		}
		cnt = ct
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
