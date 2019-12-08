package p1283

func smallestDivisor(nums []int, threshold int) int {

	check := func(divisor int) bool {
		var sum int

		for _, num := range nums {
			sum += divide(num, divisor)
		}

		return sum <= threshold
	}

	left := 1
	right := 10000000

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func divide(a, b int) int {
	c := a / b
	if c*b == a {
		return c
	}
	return c + 1
}
