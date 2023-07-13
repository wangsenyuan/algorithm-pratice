package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	A := readNInt64s(reader, n)
	res := solve(A, k)
	fmt.Println(res)
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

func solve(A []int64, k int) int64 {
	n := len(A)
	// dp[i][j][k] 表示前i个数，在保证高x位，且组成k组时的最大值

	dp := make([]uint64, n+1)
	// k <= 50

	check := func(x int64) bool {
		for i := 0; i <= n; i++ {
			dp[i] = 0
		}
		dp[0] |= 1
		for i := 1; i <= n; i++ {
			var sum int64
			for j := i; j > 0; j-- {
				sum += A[j-1]
				if sum&x == x {
					dp[i] |= dp[j-1] << 1
				}
			}
		}
		return dp[n]&(1<<k) > 0
	}

	var dfs func(x int64, d int) int64

	dfs = func(x int64, d int) int64 {
		if d < 0 {
			return x
		}

		// s1 & x == x and s2 & x == x and so one
		if check(x | int64(1)<<d) {
			return dfs(x|int64(1)<<d, d-1)
		}
		return dfs(x, d-1)
	}

	return dfs(0, 63)
}
