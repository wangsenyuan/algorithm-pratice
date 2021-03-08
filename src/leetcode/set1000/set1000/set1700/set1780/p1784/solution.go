package p1784

func minElements(nums []int, limit int, goal int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	if sum == goal {
		return 0
	}
	a := abs(sum - goal)

	return (a + limit - 1) / limit
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
