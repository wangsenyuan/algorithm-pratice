package p1780

func checkPowersOfThree(n int) bool {
	for n > 1 {
		m, r := n/3, n%3
		if r > 1 {
			return false
		}
		n = m
	}

	return n == 1
}
