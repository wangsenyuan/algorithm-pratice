package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, k, A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n, K int, A []int, B []int) int {
	// k <= n <= 40
	// max(min(A, B))
	// max(S(min(A[i], B[i])))
	X := n * 40
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = make([]int, X+1)
			for x := 0; x <= X; x++ {
				dp[i][j][x] = math.MinInt32
			}
		}
	}

	dp[0][0][0] = 0

	var sa int

	for i := 0; i < n; i++ {
		sa += A[i]
		for k := 0; k <= K; k++ {
			for x := 0; x <= sa; x++ {
				dp[i+1][k][x] = max(dp[i+1][k][x], dp[i][k][x])
				if k+1 <= K && x+A[i] <= X {
					dp[i+1][k+1][x+A[i]] = max(dp[i+1][k+1][x+A[i]], dp[i][k][x]+B[i])
				}
			}
		}
	}
	var best int
	for x := 0; x <= sa; x++ {
		y := dp[n][K][x]
		best = max(best, min(x, y))
	}
	return best

}

func solve1(n, K int, A []int, B []int) int {
	// k <= n <= 40
	// max(min(A, B))
	// A[i] <= 40
	// dp[i][j][sa][sb] = 前i个数的，由j个数组成的数, 且sum(A)的值最小是sa, sum(B)的值最小为sb，是否能达到
	// n * k * n * 40 * n * 40
	// n^^4 * 1600
	X := n * 40
	dp := make([][][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]bool, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = make([][]bool, X+1)
			for x := 0; x <= X; x++ {
				dp[i][j][x] = make([]bool, X+1)
			}
		}
	}

	dp[0][0][0][0] = true

	var sa int
	var sb int

	for i := 0; i < n; i++ {
		sa += A[i]
		sb += B[i]
		for k := 0; k <= K; k++ {
			for x := 0; x <= sa; x++ {
				for y := 0; y <= sb; y++ {
					if dp[i][k][x][y] {
						dp[i+1][k][x][y] = true
						a := A[i]
						b := B[i]
						if k < K && x+a <= X && y+b <= X {
							dp[i+1][k+1][x+a][y+b] = true
						}
					}
				}
			}
		}
	}
	var best int
	for x := 0; x <= sa; x++ {
		for y := 0; y <= sb; y++ {
			if dp[n][K][x][y] {
				best = max(best, min(x, y))
			}
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
