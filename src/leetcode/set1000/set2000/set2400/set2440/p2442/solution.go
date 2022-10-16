package p2442

func sumOfNumberAndReverse(num int) bool {
	// k + reverse(k) = num
	if num == 0 {
		return true
	}

	for x := 1; x < num; x++ {
		y := reverse(x)
		if x+y == num {
			return true
		}
	}
	return false
}

func reverse(num int) int {
	var res int

	for num > 0 {
		x := num % 10
		res = res*10 + x
		num /= 10
	}
	return res
}
