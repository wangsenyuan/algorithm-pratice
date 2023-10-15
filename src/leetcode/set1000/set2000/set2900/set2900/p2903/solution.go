package p2903

func findIndices(nums []int, indexDifference int, valueDifference int) []int {

	for j := 0; j < len(nums); j++ {
		for i := j - indexDifference; i >= 0; i-- {
			if abs(nums[j]-nums[i]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
