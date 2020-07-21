package p136

func singleNumber(nums []int) int {
	var res = 0

	for _, num := range nums {
		res ^= num
	}

	return res
}
