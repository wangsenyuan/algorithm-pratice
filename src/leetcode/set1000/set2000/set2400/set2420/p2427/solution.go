package p2427

func commonFactors(a int, b int) int {
	var res = 1
	for i := 2; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}

	return res
}
