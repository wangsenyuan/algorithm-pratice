package p2598

func findSmallestInteger(nums []int, value int) int {

	// nums[i] = k * value + nums[i] % value
	// 所以按照余数分组
	cnt := make([]int, value)
	for _, num := range nums {
		r := num % value
		if r < 0 {
			r += value
		}
		cnt[r]++
	}
	var mex int
	for {
		r := mex % value
		if cnt[r] == 0 {
			break
		}
		cnt[r]--
		mex++
	}

	return mex
}
