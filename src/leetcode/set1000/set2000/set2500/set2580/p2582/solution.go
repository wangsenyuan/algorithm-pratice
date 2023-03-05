package p2582

func passThePillow(n int, time int) int {
	if n == 1 {
		return 1
	}

	x, y := time/(n-1), time%(n-1)
	if x%2 == 0 {
		return y + 1
	}
	return n - y
}
