package p5625

func numberOfMatches(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 == 0 {
		return numberOfMatches(n / 2) + n /2
	}
	return n/2 + numberOfMatches(n/2 + 1)
}
