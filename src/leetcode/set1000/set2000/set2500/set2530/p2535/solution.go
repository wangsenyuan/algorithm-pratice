package p2535

func differenceOfSum(nums []int) int {

	var sum int
	var digit int

	for _, num := range nums {
		sum += num

		for num > 0 {
			digit += num % 10
			num /= 10
		}
	}

	return abs(sum - digit)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
