package p2348

func zeroFilledSubarray(nums []int) int64 {

	var res int64

	n := len(nums)

	for i := 0; i < n; {
		if nums[i] != 0 {
			i++
			continue
		}
		j := i
		for i < n && nums[i] == 0 {
			i++
		}
		l := int64(i - j)
		res += l * (l + 1) / 2
	}

	return res

}
