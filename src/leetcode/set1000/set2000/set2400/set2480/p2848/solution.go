package p2848

func countSubarrays(nums []int, k int) int {
	n := len(nums)

	sum := make([]int, n+1)

	var pos int

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i]
		if nums[i] > k {
			sum[i+1]++
		}
		if nums[i] == k {
			pos = i
		}
	}

	cnt_diff := make(map[int]int)
	cnt_diff[0]++

	for i := pos + 1; i < n; i++ {
		// 在[pos...i] 有a个比k大
		a := sum[i+1] - sum[pos+1]
		b := i - pos - a
		cnt_diff[a-b]++
	}
	var res int
	for i := 0; i <= pos; i++ {
		// 有x个比k大
		x := sum[pos] - sum[i]
		//有y个比k小
		y := pos - i - x
		// 假设k是最中值, y + a == x + b
		// y - x = b - a
		res += cnt_diff[y-x]
		// y + a + 1 == x + b
		res += cnt_diff[y+1-x]
	}
	return res
}
