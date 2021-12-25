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
		n, D := readTwoNums(reader)
		c := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			c[i] = readNNums(reader, n-1-i)
		}
		C := convert(n, c)
		res := solve(n, C, D)
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

func convert(n int, c [][]int) [][]int {
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		if i+1 < n {
			copy(C[i][i+1:], c[i])
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			C[j][i] = C[i][j]
		}
	}
	return C
}

const INF = 1 << 30

func solve(n int, C [][]int, D int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	dist := func(i, j int) int {
		return D*(j-i) + C[i][j]
	}
	// dp[x][y] = suppose that we have placed numbers from xx to nn and numbers on the border are xx(it will be also on the border) and yy, what’s the minimal length in this case
	dp[n-1][n-1] = 0
	for x := n - 1; x > 0; x-- {
		for y := x; y < n; y++ {
			// try to compute dp[x-1][y]
			dp[x-1][x] = min(dp[x-1][x], dp[x][y]+dist(x-1, y))
			dp[x-1][y] = min(dp[x-1][y], dp[x][y]+dist(x-1, x))
		}
	}
	var best = INF
	for i := 1; i < n; i++ {
		best = min(best, dp[0][i]+dist(0, i))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
