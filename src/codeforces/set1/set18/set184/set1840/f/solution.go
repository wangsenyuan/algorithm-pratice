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
		k := readNum(reader)
		shots := make([][]int, k)
		for i := 0; i < k; i++ {
			shots[i] = readNNums(reader, 3)
		}
		res := solve(n, m, shots)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, m int, shots [][]int) int {
	free := make([][][]bool, n+1)
	dp := make([][][]bool, n+1)
	r := len(shots)
	for i := 0; i <= n; i++ {
		free[i] = make([][]bool, m+1)
		dp[i] = make([][]bool, m+1)
		for k := 0; k <= m; k++ {
			free[i][k] = make([]bool, r+1)
			dp[i][k] = make([]bool, r+1)
			for j := 0; j <= r; j++ {
				free[i][k][j] = true
			}
		}
	}

	for i := 0; i < r; i++ {
		t, d, coord := shots[i][0], shots[i][1], shots[i][2]
		if d == 1 {
			for j := 0; j <= m; j++ {
				if 0 <= t-coord-j && t-coord-j <= r {
					free[coord][j][t-coord-j] = false
				}
			}
		} else {
			for j := 0; j <= n; j++ {
				if 0 <= t-coord-j && t-coord-j <= r {
					free[j][coord][t-coord-j] = false
				}
			}
		}
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= r; k++ {
				if i == 0 && j == 0 && k == 0 {
					dp[i][j][k] = true
				}
				if free[i][j][k] {
					if i > 0 && dp[i-1][j][k] {
						dp[i][j][k] = true
					}
					if j > 0 && dp[i][j-1][k] {
						dp[i][j][k] = true
					}
					if k > 0 && dp[i][j][k-1] {
						dp[i][j][k] = true
					}
				}
			}
		}
	}

	for k := 0; k <= r; k++ {
		if dp[n][m][k] {
			return n + m + k
		}
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
