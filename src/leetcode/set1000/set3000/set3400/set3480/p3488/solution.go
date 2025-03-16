package p3488

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	arr := make([]int, n*2)
	copy(arr, nums)
	copy(arr[n:], nums)
	pos := make(map[int]int)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i, x := range arr {
		if j, ok := pos[x]; ok {
			ans[i%n] = i - j
		}
		pos[x] = i
	}
	clear(pos)
	for i := len(arr) - 1; i >= 0; i-- {
		if j, ok := pos[arr[i]]; ok {
			if ans[i%n] == -1 || ans[i%n] > j-i {
				ans[i%n] = j - i
			}
		}
		pos[arr[i]] = i
	}
	res := make([]int, len(queries))
	for i, j := range queries {
		res[i] = ans[j]
		if res[i] >= n {
			res[i] = -1
		}
	}
	return res
}
