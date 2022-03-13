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

	price := readNNums(reader, 5)

	n := readNum(reader)

	Q := make([][]int, n)

	for i := 0; i < n; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(price, Q)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(price []int, Q [][]int) []int64 {
	beauty := []int{-2, -1, 0, 1, 2}

	const OFFSET = 200
	const INF = math.MaxInt64 >> 1

	dp := make([][]int64, 101)
	for i := 0; i <= 100; i++ {
		dp[i] = make([]int64, 401)
		for j := 0; j <= 400; j++ {
			dp[i][j] = INF
		}
	}
	// dp[0][0]
	dp[0][200] = 0

	for i := 1; i <= 100; i++ {
		for j := 0; j < 5; j++ {
			dp[i][i*beauty[j]+OFFSET] = min(dp[i][i*beauty[j]+OFFSET], int64(i)*int64(price[j]))
		}
	}

	for i := 2; i <= 100; i++ {
		for j := 0; j <= 400; j++ {
			var tmp int64 = INF
			for k := 0; k < 5; k++ {
				if j-beauty[k] >= 0 && j-beauty[k] <= 400 {
					tmp = min(tmp, dp[i-1][j-beauty[k]]+int64(price[k]))
				}
			}
			dp[i][j] = min(dp[i][j], tmp)
		}
	}

	ans := make([]int64, len(Q))
	for i, cur := range Q {
		ans[i] = dp[cur[0]][cur[1]+OFFSET]
		if ans[i] >= INF {
			ans[i] = -1
		}
	}
	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
