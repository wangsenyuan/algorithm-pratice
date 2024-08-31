package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = readNNums(reader, m)
		}
		res := solve(a)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 60

func solve(a [][]int) int {
	n := len(a)
	m := len(a[0])
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	process := func(r, c int) int {
		// r, c is fixed
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dp[i][j] = inf
			}
		}

		dp[r][c] = 0
		for i := r; i >= 0; i-- {
			for j := c; j >= 0; j-- {
				if i == r && j == c {
					continue
				}
				expect := a[r][c] - (r - i) - (c - j)

				tmp := inf
				if i+1 <= r {
					tmp = min(tmp, dp[i+1][j])
				}
				if j+1 <= c {
					tmp = min(tmp, dp[i][j+1])
				}

				if a[i][j] >= expect {
					dp[i][j] = a[i][j] - expect + tmp
				}
			}
		}

		for i := r; i < n; i++ {
			for j := c; j < m; j++ {
				if i == r && j == c {
					continue
				}
				expect := a[r][c] + (i - r) + (j - c)
				tmp := inf
				if i-1 >= r {
					tmp = min(tmp, dp[i-1][j])
				}
				if j-1 >= c {
					tmp = min(tmp, dp[i][j-1])
				}
				if a[i][j] >= expect {
					dp[i][j] = a[i][j] - expect + tmp
				}
			}
		}

		return dp[0][0] + dp[n-1][m-1]
	}

	best := inf

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			best = min(best, process(i, j))
		}
	}

	return best
}
