package p3151

func waysToReachStair(k int) int {
	if k == 0 {
		return 2
	}

	C := make([][]int, 40)

	for i := 0; i < 40; i++ {
		C[i] = make([]int, 40)
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}

	var res int
	cur := 1

	jumps := -1

	for {
		if cur >= k {
			x := cur - k
			if x <= jumps+2 {
				res += C[jumps+2][x]
			} else {
				break
			}
		}
		jumps++
		cur += 1 << jumps
	}

	return res
}
