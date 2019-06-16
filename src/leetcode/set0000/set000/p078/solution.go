package p078

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	last := nums[len(nums)-1]
	result := subsets(nums[:len(nums)-1])
	n := len(result)

	for i := 0; i < n; i++ {
		tmp := make([]int, len(result[i]))
		copy(tmp, result[i])
		result = append(result, append(tmp, last))
	}
	return result
}
