package p2750

func makeTheIntegerZero(num1 int, num2 int) int {
	for x := 1; ; x++ {
		num := int64(num1) - int64(x)*int64(num2)
		if num < int64(x) {
			return -1
		}
		if digit_count(num) <= x {
			return x
		}
	}
}

func mst(num int64) int {
	var h int = -1
	for num > 0 {
		h++
		num >>= 1
	}
	return h
}

func digit_count(num int64) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
