package lcp02

func fraction(cont []int) []int {
	n := len(cont)
	res := []int{cont[n-1], 1}

	for i := n - 2; i >= 0; i-- {
		// a[i] + 1 / res
		res[0], res[1] = res[1], res[0]
		res[0] += cont[i] * res[1]
		g := gcd(res[0], res[1])
		res[0] /= g
		res[1] /= g
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
