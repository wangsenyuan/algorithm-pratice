package p678

func checkValidString1(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n+1)
		dp[i][i] = true
		if s[i] == '*' {
			dp[i][i+1] = true
		}
	}

	for k := 1; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			j := i + k
			if !dp[i][j] {
				if s[i] == '(' {
					if s[j-1] == ')' {
						dp[i][j] = dp[i+1][j-1]
					} else if s[j-1] == '*' {
						dp[i][j] = dp[i+1][j-1] || dp[i][j-1]
					}

				} else if s[i] == '*' {
					if s[j-1] == ')' {
						dp[i][j] = dp[i+1][j-1] || dp[i+1][j]
					} else if s[j-1] == '*' {
						dp[i][j] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j-1]
					}
				}
			}

			if !dp[i][j] && s[i] != ')' && s[j-1] != '(' {
				for a := i + 1; a < j; a++ {
					if dp[i][a] && dp[a][j] {
						dp[i][j] = true
						break
					}
				}
			}
		}
	}
	return dp[0][n]
}

func checkValidString(s string) bool {
	n := len(s)
	var low, high int

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			low++
		} else {
			low--
		}

		if s[i] != ')' {
			high++
		} else {
			high--
		}
		if high < 0 {
			return false
		}
		if low < 0 {
			low = 0
		}
	}
	return low == 0
}
func checkValidString2(s string) bool {
	n := len(s)

	var open, star int

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			open++
		} else if s[i] == ')' {
			open--
		} else {
			star++
		}
		if open < 0 {
			if star == 0 {
				return false
			}
			open++
			star--
		}
	}

	open = 0
	star = 0

	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			open--
		} else if s[i] == ')' {
			open++
		} else {
			star++
		}
		if open < 0 {
			if star == 0 {
				return false
			}
			open++
			star--
		}
	}

	return true
}
