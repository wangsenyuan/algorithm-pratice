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
		n := readNum(reader)
		A := readNNums(reader, n)

		res := solve(A)

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

func solve(A []int) int64 {
	n := len(A)
	dp := make([][]int64, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
		dp[i][i] = math.MinInt64
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			k := max2(i, j) + 1
			if k > n {
				continue
			}
			// try calculate dp[i][k] and dp[k][j]
			if i == 0 {
				dp[k][j] = max(dp[k][j], dp[i][j])
			} else {
				dp[k][j] = max(dp[k][j], dp[i][j]+int64(abs(A[i-1]-A[k-1])))
			}
			if j == 0 {
				dp[i][k] = max(dp[i][k], dp[i][j])
			} else {
				dp[i][k] = max(dp[i][k], dp[i][j]+int64(abs(A[j-1]-A[k-1])))
			}
		}
	}

	var res int64

	for i := 0; i <= n; i++ {
		res = max(res, dp[n][i])
		res = max(res, dp[i][n])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
