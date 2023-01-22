package p2543

func isReachable(x int, y int) bool {

	for x&1 == 0 {
		x >>= 1
	}

	for y&1 == 0 {
		y >>= 1
	}

	for x != 1 && y != 1 {
		if x > y {
			x, y = y, x
		}
		y %= x
		if y == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
