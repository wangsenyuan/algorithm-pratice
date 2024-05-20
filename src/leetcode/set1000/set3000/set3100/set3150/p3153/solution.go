package p3153

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	// b[i] = 1, if b[i] % 2 != b[i+1] % 2
	b := make([]int, n)
	for i := 1; i < n; i++ {
		b[i] = b[i-1]
		if nums[i-1]&1 != nums[i]&1 {
			b[i]++
		}

	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = b[r]-b[l] == r-l
	}

	return ans
}
