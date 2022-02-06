package p2165

import "sort"

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}
	var sign int64 = 1
	if num < 0 {
		sign = -1
		num *= -1
	}
	var digits []int

	for num > 0 {
		digits = append(digits, int(num%10))
		num /= 10
	}
	sort.Ints(digits)
	if sign > 0 {
		var i int
		for digits[i] == 0 {
			i++
		}
		digits[0], digits[i] = digits[i], digits[0]
		var res int64
		for i := 0; i < len(digits); i++ {
			res = res*10 + int64(digits[i])
		}

		return res
	}
	var res int64

	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + int64(digits[i])
	}

	return res * -1
}
