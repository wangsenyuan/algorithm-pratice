package p997

func findJudge(N int, trust [][]int) int {
	//trust count
	a := make([]int, N)
	//been trusted
	b := make([]int, N)

	for _, x := range trust {
		i, j := x[0], x[1]
		a[i-1]++
		b[j-1]++
	}

	ans := -1
	for i := 0; i < N; i++ {
		if a[i] == 0 {
			//property one
			if b[i] == N-1 {
				//property two
				if ans == -1 {
					ans = i + 1
				} else {
					//property three not hold
					return -1
				}
			}
		}
	}

	return ans
}
