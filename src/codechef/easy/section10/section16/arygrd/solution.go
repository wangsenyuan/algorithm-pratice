package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	line := readNNums(scanner, 3)
	m, n, x := line[0], line[1], line[2]
	grid := make([]string, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		grid[i] = scanner.Text()
	}

	fmt.Println(solve(m, n, x, grid))
}

func solve(m, n, x int, grid []string) int {

	dp := make([][][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, x+1)
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			for p := 0; p <= x; p++ {
				var cnt1 int
				temp := -1
				for p1 := j; p1 <= min(j+x, n-1); p1++ {
					if grid[i][p1] == '#' {
						break
					}
					if grid[i][p1] == '*' {
						cnt1++
					}
					var cnt2 int
					for p2 := max(p1, j+p); p2 <= min(j+x, n-1); p2++ {
						if grid[i][p2] == '#' {
							break
						}
						if grid[i][p2] == '*' && p2 > p1 {
							cnt2++
						}
						if (i < m-1 || (p1 == n-1 && p2 == n-1)) && dp[i+1][p1][p2-p1] != -1 {
							temp = max(temp, dp[i+1][p1][p2-p1]+cnt1+cnt2)
						}
					}
				}
				dp[i][j][p] = temp
			}
		}
	}

	return dp[0][0][0]
}

func solve2(m, n, x int, grid []string) int {
	if grid[0][0] == '#' {
		return -1
	}

	dp := make([][][]int, m)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, x+1)
			for k := 0; k <= x; k++ {
				dp[i][j][k] = -2
			}
		}
	}

	var dfs func(i, j, k int) int

	dfs = func(i, j, k int) int {
		if i == m-1 {
			if j+x < n-1 {
				return -1
			}

			var c int
			var u = 0
			for j+u < n {
				if grid[i][j+u] == '#' {
					return -1
				}
				if grid[i][j+u] == '*' {
					c++
				}
				u++
			}

			return c
		}
		if dp[i][j][k] >= -1 {
			return dp[i][j][k]
		}

		best := -1

		var c1 int

		// K := j + k

		for a := j; a <= j+x && a < n; a++ {
			if grid[i][a] == '#' {
				break
			}

			if grid[i][a] == '*' {
				c1++
			}

			b := max(j+k, a)
			var c2 int

			for b <= j+x && b < n && grid[i][b] != '#' {
				if grid[i][b] == '*' && b > a {
					c2++
				}

				if grid[i+1][a] != '#' && grid[i+1][b] != '#' {
					tmp := dfs(i+1, a, b-a)
					if tmp >= 0 {
						best = max(tmp+c1+c2, best)
					}
				}

				b++
			}
		}

		dp[i][j][k] = best

		return best
	}

	res := dfs(0, 0, 0)

	if res < 0 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(m, n, x int, grid []string) int {
	if grid[0][0] == '#' {
		return -1
	}

	// x = min(n, x)

	// dp[i][j][u][v][w] = twos players at (i, j), (i, j + u), first move from left v, and second move from left w

	dp := make([][][][][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][][][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([][][]int, x+1)
			for u := 0; u <= x; u++ {
				dp[i][j][u] = make([][]int, x+1)
				for v := 0; v <= x; v++ {
					dp[i][j][u][v] = make([]int, x+1)
					for w := 0; w <= x; w++ {
						dp[i][j][u][v][w] = -1
					}
				}
			}
		}
	}

	dp[0][0][0][0][0] = 0
	if grid[0][0] == '*' {
		dp[0][0][0][0][0] = 1
	}

	// dp[i][j][u][v][w] = twos players at (i, j), (i, j + u),
	// first move from left v, and second move from left w

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for u := 0; u <= x; u++ {
				for v := 0; v <= x; v++ {
					for w := 0; w <= x; w++ {
						if dp[i][j][u][v][w] < 0 {
							continue
						}
						c1 := j - v
						j1 := j
						j2 := j + u
						c2 := j2 - w
						if v < x && j1+1 <= c2 {
							// move first one right
							if grid[i][j1+1] != '#' {
								// not stone
								var star int
								if j1+1 < c2 && grid[i][j1+1] == '*' {
									star++
								}
								setIfMax(&dp[i][j1+1][u-1][v+1][w], dp[i][j][u][v][w]+star)
							}
						}

						if w < x && j2 < n-1 && j2 < c1+x {
							// move second one right
							if grid[i][j2+1] != '#' {
								var star int
								if grid[i][j2+1] == '*' {
									star++
								}
								setIfMax(&dp[i][j1][u+1][v][w+1], dp[i][j][u][v][w]+star)
							}
						}
						if i < m-1 {
							// move both down one row
							if grid[i+1][j1] != '#' && grid[i+1][j2] != '#' {
								var star int
								if grid[i+1][j1] == '*' {
									star++
								}
								if grid[i+1][j2] == '*' && j1 != j2 {
									star++
								}
								setIfMax(&dp[i+1][j1][u][0][0], dp[i][j][u][v][w]+star)
							}
						}
					}
				}
			}
		}
	}

	res := -1

	for u := 0; u <= x; u++ {
		j := n - 1 - u
		if j < 0 {
			continue
		}
		for v := 0; v <= x; v++ {
			for w := 0; w <= x; w++ {
				setIfMax(&res, dp[m-1][j][u][v][w])
			}
		}
	}

	return res
}

func setIfMax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
