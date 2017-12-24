package p747

func dominantIndex(nums []int) int {
	var i int

	for j := 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			i = j
		}
	}

	for j := 0; j < len(nums); j++ {
		if i == j {
			continue
		}
		if 2*nums[j] > nums[i] {
			return -1
		}
	}
	return i
}
