package p728

func selfDividingNumbers(left int, right int) []int {
	n := right - left + 1
	res := make([]int, 0, n)

	for x := left; x <= right; x++ {
		if selfDividing(x) {
			res = append(res, x)
		}
	}

	return res
}

func selfDividing(num int) bool {
	tmp := num
	for tmp > 0 {
		digit := tmp % 10
		if digit == 0 || num%digit != 0 {
			return false
		}
		tmp = tmp / 10
	}
	return true
}
