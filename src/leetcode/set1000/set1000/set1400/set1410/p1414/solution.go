package p1414

func findMinFibonacciNumbers(k int) int {
	// F[i] = F[i-1] + F[i-2] = F[i-2] + F[i-3] + F[i-2]
	F := make([]int, 0, 10)

	F = append(F, 1)
	F = append(F, 1)

	for {
		n := len(F)
		x := F[n-1] + F[n-2]
		if x > k || x < 0 {
			break
		}
		F = append(F, x)
	}

	var res int

	i := len(F) - 1
	var y int

	for y < k {
		if y+F[i] <= k {
			y += F[i]
			res++
			continue
		}
		i--
	}

	return res
}
