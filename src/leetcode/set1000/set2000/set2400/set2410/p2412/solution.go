package p2412

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
