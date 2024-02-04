package p3029

func returnToBoundaryCount(nums []int) int {
	var res int

	var sum int

	for _, num := range nums {
		sum += num
		if sum == 0 {
			res++
		}
	}

	return res
}
