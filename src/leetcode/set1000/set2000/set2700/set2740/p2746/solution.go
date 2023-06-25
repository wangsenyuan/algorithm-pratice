package p2746

const inf = 1 << 30

func longestString(x int, y int, z int) int {
	// AA 后面只能更BB
	// BB后面只能跟AA，或者AB
	// AB后面只能跟AA，或者AB
	// 如果在某个位置放置了AA，那么只能在它后面放置BB
	i := min(x, y)
	a := x - i
	b := y - i
	// a个AABB变成了AB
	//c := i + z
	var res int
	if a == 0 {
		// 只有BB和AB, 只时候只能放置一个BB在最前面，其他的全是AB
		res = 2*i + z
		if b > 0 {
			res++
		}
	} else {
		// b == 0
		// AA只能放置在最后面
		res = 2*i + z + 1
	}

	return 2 * res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestString1(x int, y int, z int) int {
	// x AA, y BB, z AB
	// 不能有AAA, BBB
	// 所以就是不能有连续的两个 AA AA, BB, BB 或者 AA AB , AB BB
	// 但是可以有连续的 ABAB
	// AA 后面只能跟 BB, BB后面可以跟AA或者AB， AB后面只能跟AB, AA
	// dp[a][b][c][x] 表示使用了第一种a个，第二种b个，第三种c个时，且最后一个时x（0， 1， 1）时的最大长度
	dp := make([][][][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([][][]int, y+1)
		for j := 0; j <= y; j++ {
			dp[i][j] = make([][]int, z+1)
			for k := 0; k <= z; k++ {
				dp[i][j][k] = make([]int, 3)
				for u := 0; u < 3; u++ {
					dp[i][j][k][u] = -inf
				}
			}
		}
	}
	dp[1][0][0][0] = 1
	dp[0][1][0][1] = 1
	dp[0][0][1][2] = 1

	var best int

	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			for k := 0; k <= z; k++ {
				for u := 0; u < 3; u++ {
					if dp[i][j][k][u] <= 0 {
						continue
					}
					best = max(best, dp[i][j][k][u])
					if u == 0 {
						// AA => BB
						if j+1 <= y {
							dp[i][j+1][k][1] = max(dp[i][j+1][k][1], dp[i][j][k][u]+1)
						}
					} else {
						// BB => AB or AA
						// AB => AB or AA
						if i+1 <= x {
							dp[i+1][j][k][0] = max(dp[i+1][j][k][0], dp[i][j][k][u]+1)
						}
						if k+1 <= z {
							dp[i][j][k+1][2] = max(dp[i][j][k+1][2], dp[i][j][k][u]+1)
						}
					}
				}
			}
		}
	}
	return best * 2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
