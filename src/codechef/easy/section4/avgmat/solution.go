package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		g := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			g[i] = make([]byte, m)
			copy(g[i], scanner.Bytes())
		}

		res := solve(n, m, g)

		for i := 0; i < len(res); i++ {
			if i < len(res)-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}

	}
}

func solve(n, m int, g [][]byte) []int64 {

	lr := make([][]int64, n)
	rl := make([][]int64, n)
	for i := 0; i < n; i++ {
		lr[i] = make([]int64, m)
		rl[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			if g[i][j] == '1' {
				lr[i][j] = 1
				rl[i][j] = 1
			}

			if i > 0 && j > 0 {
				lr[i][j] += lr[i-1][j-1]
			}

			if i > 0 && j < m-1 {
				rl[i][j] += rl[i-1][j+1]
			}
		}
	}
	L := n + m - 2

	getLR := func(x, y int) int64 {
		if y-x >= m || x-y >= n {
			return 0
		}
		if min(x, y) < 0 {
			return 0
		}

		if x >= n || y >= m {
			d := max(x-n+1, y-m+1)
			x -= d
			y -= d
		}
		return lr[x][y]
	}

	getRL := func(x, y int) int64 {
		if x+y < 0 || x+y > L {
			return 0
		}
		if x < 0 || y > m {
			return 0
		}
		if y < 0 {
			x += y
			y = 0
		} else if x >= n {
			y += x - n + 1
			x = n - 1
		}
		return rl[x][y]

	}

	res := make([]int64, L)

	for l := 1; l <= L; l++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if g[i][j] == '0' {
					continue
				}

				var cnt int64

				cnt += getRL(i, j-l) - getRL(i-l-1, j+1)
				cnt += getRL(i+l, j) - getRL(-1, j+l+1)
				cnt += getLR(i, j+l) - getLR(i-l-1, j-1)
				cnt += getLR(i+l, j) - getLR(i-1, j-l-1)
				if i-l >= 0 && g[i-l][j] == '1' {
					cnt--
				}
				if j-l >= 0 && g[i][j-l] == '1' {
					cnt--
				}
				if i+l < n && g[i+l][j] == '1' {
					cnt--
				}
				if j+l < m && g[i][j+l] == '1' {
					cnt--
				}
				res[l-1] += cnt
			}
		}
	}

	for i := 0; i < L; i++ {
		res[i] /= 2
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
