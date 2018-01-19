package p029

import "math"

func divide(dividend int, divisor int) int {
	dvd, dvs := int64(dividend), int64(divisor)

	positive := (dvd > 0 && dvs > 0) || (dvd < 0 && dvs < 0)

	if dvd < 0 {
		dvd *= -1
	}

	if dvs < 0 {
		dvs *= -1
	}
	var res int64

	for dvd >= dvs {
		tmp, times := dvs, int64(1)

		for dvd >= (tmp << 1) {
			tmp <<= 1
			times <<= 1
		}
		res += times
		dvd -= tmp
	}

	if !positive {
		res *= -1
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return int(res)
}
