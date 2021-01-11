package main

func totalMoney(n int) int {
	var res int
	cur := 1
	for n > 0 {
		r := min(7, n)
		res += (cur + cur + r - 1) * r / 2
		cur++
		n -= 7
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
