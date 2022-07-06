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
		D, n, m := readThreeNums(reader)
		B := make([][]int, m)
		for i := 0; i < m; i++ {
			B[i] = readNNums(reader, n)
		}
		C := make([][]int, m)
		for i := 0; i < m; i++ {
			C[i] = readNNums(reader, m)
		}
		res := solve(D, B, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const INF = 1 << 30

func solve(D int, B [][]int, C [][]int) int {
	m := len(B)
	n := len(B[0])

	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cost[i][j] = C[i][j]
		}
		cost[i][i] = 0
	}

	for k := 0; k < m; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k][j])
			}
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = INF
		}
	}

	for j := 0; j < m; j++ {
		dp[0][j] = B[j][0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+cost[k][j])
			}
			dp[i][j] += B[j][i]
		}
	}
	res := INF

	for j := 0; j < m; j++ {
		res = min(res, dp[n-1][j])
	}

	if res <= D {
		return D - res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
