package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := make([]string, n)
		for i := 0; i < n; i++ {
			A[i], _ = reader.ReadString('\n')
			A[i] = A[i][:m]
		}
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

const X = 256

/**
dp[x][y] = 高8位和x相同，但低8位与y不同seq的个数
*/
var dp [][]int64

func init() {
	dp = make([][]int64, X)
	for i := 0; i < X; i++ {
		dp[i] = make([]int64, X)
	}
}

func reset() {
	for i := 0; i < X; i++ {
		for j := 0; j < X; j++ {
			dp[i][j] = 0
		}
	}
}

func solve(n int, A []string) int {
	reset()
	var ans int64
	for i := 0; i < n; i++ {
		x := parseAsNum(A[i])
		var ways int64 = 1

		for j := 0; j < X; j++ {
			if (x>>8)&j == 0 {
				add(&ways, dp[j][x&255])
			}
		}

		add(&ans, ways)

		for j := 0; j < X; j++ {
			if x&j == 0 {
				add(&dp[x>>8][j], ways)
			}
		}

	}

	return int(ans)
}

func add(a *int64, b int64) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func parseAsNum(s string) int {
	var res int
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			res |= 1 << uint(n-1-i)
		}
	}

	return res
}
