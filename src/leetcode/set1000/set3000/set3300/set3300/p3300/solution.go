package p3300

func minElement(nums []int) int {
	best := 1 << 30

	for _, num := range nums {
		best = min(num, digitSum(num))
	}

	return best
}

func digitSum(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
