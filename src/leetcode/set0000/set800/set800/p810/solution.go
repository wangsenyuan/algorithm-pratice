package p810

func xorGame(nums []int) bool {

	var res int

	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	if res == 0 {
		return true
	}

	return len(nums)%2 == 0
}
