package p2295

func arrayChange(nums []int, operations [][]int) []int {
	pos := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		pos[nums[i]] = i
	}

	for _, op := range operations {
		x, y := op[0], op[1]
		i := pos[x]
		delete(pos, x)
		pos[y] = i
	}

	res := make([]int, n)
	for k, i := range pos {
		res[i] = k
	}
	return res
}
