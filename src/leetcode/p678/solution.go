package main

import "fmt"

func checkValidString(s string) bool {
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

func main() {
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("())"))
	fmt.Println(checkValidString("*"))

}
