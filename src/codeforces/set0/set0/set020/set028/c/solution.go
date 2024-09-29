package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, m)
	res := solve(n, m, a)

	fmt.Printf("%.9f\n", res)
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

const N = 60

var C [N][N]int

func init() {
	C[0][0] = 1
	for i := 1; i < N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
}

func solve(n int, m int, a []int) float64 {
	// len(a) = m
	var mq int

	for _, x := range a {
		mq = max(mq, (n+x-1)/x)
	}

	vis := make([][][]bool, n+1)
	dp := make([][][]float64, n+1)
	for i := 0; i <= n; i++ {
		vis[i] = make([][]bool, m+1)
		dp[i] = make([][]float64, m+1)
		for j := 0; j <= m; j++ {
			vis[i][j] = make([]bool, mq+1)
			dp[i][j] = make([]float64, mq+1)
		}
	}

	calc := func(i int, j int, c int) float64 {
		x := math.Pow(1/float64(j), float64(c))
		y := math.Pow(float64(j-1)/float64(j), float64(i-c))
		z := float64(C[i][c])

		return x * y * z
	}

	var dfs func(i int, j int, k int) float64

	dfs = func(i int, j int, k int) float64 {
		if i == 0 {
			return float64(k)
		}
		if j == 0 {
			return 0
		}
		if vis[i][j][k] {
			return dp[i][j][k]
		}
		vis[i][j][k] = true
		for c := 0; c <= i; c++ {
			tmp := dfs(i-c, j-1, max(k, (c+a[j-1]-1)/a[j-1]))
			dp[i][j][k] += calc(i, j, c) * tmp
		}
		return dp[i][j][k]
	}

	return dfs(n, m, 0)
}
