package p2719

const MOD = 1000000007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	res := count_lt(num2, min_sum, max_sum)
	res = sub(res, count_lt(num1, min_sum, max_sum))
	var s int
	for i := 0; i < len(num1); i++ {
		s += int(num1[i] - '0')
	}
	if s >= min_sum && s <= max_sum {
		res = add(res, 1)
	}
	return res
}

func count_lt(num string, min_sum int, max_sum int) int {
	n := len(num)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, max_sum+1)
		}
	}
	// dp[i][j][k] = 前i个数，和num相等，且digit_sum=k时有多少个数
	dp[0][1][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			for s := 0; s <= max_sum; s++ {
				var x, y = 0, 10
				if j == 1 {
					y = int(num[i]-'0') + 1
				}
				for k := x; k < y; k++ {
					ns := s + k
					if ns > max_sum {
						break
					}
					// ns <= max_sum
					nj := 0
					if j == 1 && k == int(num[i]-'0') {
						nj = 1
					}
					dp[i+1][nj][ns] = add(dp[i+1][nj][ns], dp[i][j][s])
				}
			}
		}
	}
	var res int

	for x := min_sum; x <= max_sum; x++ {
		for j := 0; j < 2; j++ {
			res = add(res, dp[n][j][x])
		}
	}
	return res
}
