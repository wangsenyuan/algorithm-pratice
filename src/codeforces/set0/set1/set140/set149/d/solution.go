package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(s string) int {
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, n)
	var top int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			j := stack[top-1]
			top--
			pair[i] = j
			pair[j] = i
		}
	}

	merge := func(x [3][3]int, y [3][3]int) [3][3]int {
		var z [3][3]int

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				for u := 0; u < 3; u++ {
					if j > 0 && j == u {
						continue
					}
					for v := 0; v < 3; v++ {
						z[i][v] = add(z[i][v], mul(x[i][j], y[u][v]))
					}
				}
			}
		}

		return z
	}

	var dfs func(l int, r int) [3][3]int
	// dp[x][y] = (l..r)的状态为x, y时的计数
	// 0 for black, 1 for red, 2 for blue
	dfs = func(l int, r int) (res [3][3]int) {
		if l+1 == r {
			res[0][1] = 1
			res[1][0] = 1
			res[0][2] = 1
			res[2][0] = 1
			return
		}
		var cur [3][3]int
		for i := l + 1; i < r; i = pair[i] + 1 {
			tmp := dfs(i, pair[i])
			if i == l+1 {
				cur = tmp
			} else {
				// merge
				cur = merge(cur, tmp)
			}
		}
		for c1 := 0; c1 <= 2; c1++ {
			for c2 := 0; c2 <= 2; c2++ {
				for c := 1; c <= 2; c++ {
					if c != c1 {
						res[c][0] = add(res[c][0], cur[c1][c2])
					}
					if c != c2 {
						res[0][c] = add(res[0][c], cur[c1][c2])
					}
				}
			}
		}
		return
	}
	var res [3][3]int
	for i := 0; i < n; i = pair[i] + 1 {
		tmp := dfs(i, pair[i])
		if i == 0 {
			res = tmp
		} else {
			res = merge(res, tmp)
		}
	}

	var ans int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ans = add(ans, res[i][j])
		}
	}

	return ans
}
