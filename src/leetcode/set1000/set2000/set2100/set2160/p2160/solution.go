package p2160

import "sort"

func minimumSum(num int) int {
	digits := make([]int, 4)
	var i int
	for num > 0 {
		digits[i] = num % 10
		i++
		num /= 10
	}

	sort.Ints(digits)
	a := digits[0]*10 + digits[2]
	b := digits[1]*10 + digits[3]
	return a + b
}
