package p2123

func isSameAfterReversals1(num int) bool {
	tmp := int64(num)

	return reverse(reverse(tmp)) == tmp
}

func reverse(num int64) int64 {
	var res int64
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	if num%10 == 0 {
		return false
	}
	return true
}
