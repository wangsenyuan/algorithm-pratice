package p283

func moveZeroes(nums []int) {
	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if i > p {
			nums[p] = nums[i]
			nums[i] = 0
		}
		p++
	}
}
