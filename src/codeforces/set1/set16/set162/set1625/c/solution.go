package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, L, K := readThreeNums(reader)

	X := readNNums(reader, N)
	A := readNNums(reader, N)

	res := solve(L, N, K, X, A)

	fmt.Println(res)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1 << 60

func solve(L int, N int, K int, X []int, signs []int) int64 {
	// dp[i][x] = 到达位置i，且删除x个路牌时所花的最小时间
	dp := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 1; i < N; i++ {
		for j := i - 1; j >= 0; j-- {
			for x := j; x >= 0; x-- {
				// dp[j][x] = 到达位置j，且删除x个路牌的情况
				if x+(i-j-1) <= K {
					dp[i][x+(i-j-1)] = min(dp[i][x+(i-j-1)], dp[j][x]+int64(X[i]-X[j])*int64(signs[j]))
				}
			}
		}
	}
	var best int64 = INF

	for i := 0; i < N; i++ {
		for x := i; x >= 0; x-- {
			// dp[i][x]
			y := x + (N - (i + 1))
			if y <= K {
				// remove at most K
				tmp := dp[i][x] + int64(L-X[i])*int64(signs[i])
				best = min(best, tmp)
			}
		}
	}
	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
