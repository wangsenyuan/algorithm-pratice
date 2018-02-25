package p788

func rotatedDigits(N int) int {

	rotate := func(num int) (bool, int) {
		var res int

		base := 1

		for num > 0 {
			x := num % 10

			y := x
			if x == 0 || x == 1 || x == 8 {
				y = x
			} else if x == 2 {
				y = 5
			} else if x == 5 {
				y = 2
			} else if x == 6 {
				y = 9
			} else if x == 9 {
				y = 6
			} else {
				return false, 0
			}

			num /= 10
			res = y*base + res
			base *= 10
		}

		return true, res
	}

	var ans int

	for i := 2; i <= N; i++ {
		can, tmp := rotate(i)
		if can && tmp != i {
			ans++
		}
	}

	return ans
}
