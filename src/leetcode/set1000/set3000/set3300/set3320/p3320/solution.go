package p3320

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func countWinningSequences(s string) int {
	n := len(s)
	dp := make([][]int, 2*n+1)
	fp := make([][]int, 2*n+1)
	for i := 0; i <= 2*n; i++ {
		dp[i] = make([]int, 3)
		fp[i] = make([]int, 3)
	}
	// FWE
	id := map[byte]int{'F': 0, 'W': 1, 'E': 2}
	x := id[s[0]]
	for i := 0; i < 3; i++ {
		if (x+1)%3 == i {
			dp[n+1][i] = 1
		} else if i == x {
			dp[n][i] = 1
		} else {
			dp[n-1][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		x := id[s[i]]
		for j := -i; j <= i; j++ {
			for y := 0; y < 3; y++ {
				// y是前一个取值
				for z := 0; z < 3; z++ {
					if y == z {
						// not allowed
						continue
					}
					if (x+1)%3 == z {
						fp[n+j+1][z] = add(fp[n+j+1][z], dp[n+j][y])
					} else if x == z {
						fp[n+j][z] = add(fp[n+j][z], dp[n+j][y])
					} else {
						fp[n+j-1][z] = add(fp[n+j-1][z], dp[n+j][y])
					}
				}
			}
		}
		for j := 0; j <= 2*n; j++ {
			for k := 0; k < 3; k++ {
				dp[j][k] = fp[j][k]
				fp[j][k] = 0
			}
		}
	}

	var res int

	for i := n + 1; i <= 2*n; i++ {
		for j := 0; j < 3; j++ {
			res = add(res, dp[i][j])
		}
	}

	return res
}
