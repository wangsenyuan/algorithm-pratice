package p991

func brokenCalc(X int, Y int) int {
	var ans int
	for Y > X {
		ans++
		if Y&1 == 0 {
			Y >>= 1
		} else {
			Y++
		}
	}

	return ans + X - Y
}
