package p172

func trailingZeroes(n int) int {
	var ans int

	for n >= 5 {
		ans += n / 5
		n /= 5
	}

	return ans
}
