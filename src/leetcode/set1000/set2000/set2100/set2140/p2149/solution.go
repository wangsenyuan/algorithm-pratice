package p2149

func rearrangeArray(nums []int) []int {
	N := len(nums)
	n := N / 2
	pos := make([]int, n)
	neg := make([]int, n)

	for i, j, k := 0, 0, 0; i < N; i++ {
		num := nums[i]
		if num < 0 {
			neg[j] = num
			j++
		} else {
			pos[k] = num
			k++
		}
	}
	res := make([]int, N)

	for i := 0; i < N; i += 2 {
		res[i] = pos[i/2]
		res[i+1] = neg[i/2]
	}

	return res
}
