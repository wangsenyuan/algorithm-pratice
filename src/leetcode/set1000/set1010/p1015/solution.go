package p1015

var f [32]int

func init() {
	f[1] = 10
	f[2] = 9 * 9
	f[3] = f[2] * 8
	f[4] = f[3] * 7
	f[5] = f[4] * 6
	f[6] = f[5] * 5
	f[7] = f[6] * 4
	f[8] = f[7] * 3
	f[9] = f[8] * 2
	f[10] = f[9] * 1
	// f[i] = 0 when i > 10
}

func numDupDigitsAtMostN(n int) int {
	if n < 10 {
		return 0
	}

	N := int64(n)
	var cnt int
	// get count of unique digits number that have 1, 2, 3, .. digits, that less than N
	base := int64(10)
	d := 1
	for base <= N {
		cnt += f[d]
		d++
		base *= 10
	}

	// base > N
	base /= 10
	d--

	var loop func(d int, b int, lt bool, left int, mask int, num int) int

	loop = func(d int, b int, lt bool, left int, mask int, num int) int {
		if d < 0 || b == 0 {
			return 1
		}

		if lt {
			res := 1
			for d >= 0 {
				res *= left
				left--
				d--
			}
			return res
		}
		// same before digit d
		x := num / b

		var y = 1
		if mask&1 == 0 && num != n {
			// first digit can't be zero
			y = 0
		}

		var res int
		for y <= x {
			if mask&(1<<uint(y)) == 0 {
				// not used before
				res += loop(d-1, b/10, y < x, left-1, mask|(1<<uint(y)), n%b)
			}
			y++
		}

		return res
	}

	cnt2 := loop(d, int(base), false, 10, 0, n)

	return n + 1 - cnt - cnt2
}
