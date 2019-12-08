package p1281

func subtractProductAndSum(n int) int {
	var sum int64
	var product int64 = 1

	for n > 0 {
		r := n % 10
		sum += int64(r)
		product *= int64(r)
		n /= 10
	}

	return int(product - sum)
}
