package p2270

func waysToSplitArray(nums []int) int {

	var sum int64

	for _, num := range nums {
		sum += int64(num)
	}

	var pref int64

	var res int

	if pref >= sum {
		// index 0
		res--
	}

	for i := 0; i < len(nums); i++ {
		if pref >= sum {
			res++
		}
		pref += int64(nums[i])
		sum -= int64(nums[i])
	}

	return res
}
