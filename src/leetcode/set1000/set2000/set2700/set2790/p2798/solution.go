package p2798

const mod = 1e9 + 7

func add(a *int, b int) {
	*a += b
	if *a >= mod {
		*a -= mod
	}
}

func sub(a *int, b int) {
	add(a, mod-b)
}

func countSteppingNumbers(low string, high string) int {
	ans := count(high)
	sub(&ans, count(low))
	if checkStepping(low) {
		add(&ans, 1)
	}
	return ans
}

func checkStepping(num string) bool {
	n := len(num)
	if n == 1 {
		return true
	}
	prev := int(num[0] - '0')
	for i := 1; i < n; i++ {
		cur := int(num[i] - '0')
		if abs(cur-prev) != 1 {
			return false
		}
		prev = cur
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func count(num string) int {
	n := len(num)
	if n == 1 {
		return int(num[0] - '0')
	}
	// n > 1
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 10)
		}
	}
	// 还少一个0000的情况
	// dp[p][i][j][k] = i = 0 表示已经不是前导0， 1表示到目前为止还是0
	//  j 表示 0表示 < num, 1 表示 = num
	//  k 表示前一个数字时几，只有到i=0时有效
	for i := 1; i < 10; i++ {
		x := int(num[0] - '0')
		if i < x {
			dp[0][0][i] = 1
		} else {
			dp[0][1][i] = 1
		}
		if i == x {
			break
		}
	}
	for i := 1; i < n; i++ {
		p := i & 1
		for j := 0; j < 2; j++ {
			for x := 0; x < 10; x++ {
				dp[p][j][x] = 0
			}
		}
		cur := int(num[i] - '0')
		for j := 0; j < 2; j++ {
			for k := 0; k < 10; k++ {
				for _, d := range []int{-1, 1} {
					if k+d < 0 || k+d >= 10 {
						continue
					}
					if j == 0 || k+d < cur {
						add(&dp[p][0][k+d], dp[p^1][j][k])
					} else if j == 1 && k+d == cur {
						add(&dp[p][1][k+d], dp[p^1][j][k])
					}
				}
			}
		}
		for x := 1; x < 10; x++ {
			add(&dp[p][0][x], 1)
		}
	}
	var res int
	for a := 0; a < 10; a++ {
		add(&res, dp[(n-1)&1][0][a])
		add(&res, dp[(n-1)&1][1][a])
	}

	return res
}
