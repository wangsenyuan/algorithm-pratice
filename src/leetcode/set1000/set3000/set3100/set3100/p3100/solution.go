package p3100

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var ds int
	for tmp := x; tmp > 0; tmp /= 10 {
		ds += tmp % 10
	}

	if x%ds == 0 {
		return ds
	}
	return -1
}
