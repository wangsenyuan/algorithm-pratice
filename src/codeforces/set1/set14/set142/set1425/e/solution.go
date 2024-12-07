package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	d := readNNums(reader, n)
	res := solve(k, a, d)

	fmt.Println(res)
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

func solve(K int, A []int, D []int) int {
	n := len(A)

	if K == 0 {
		var suf int
		var best int
		for i := n - 1; i >= 0; i-- {
			suf += A[i]
			best = max(best, suf-D[i])
		}
		return best
	}

	if K == 1 {
		return solve1(A, D)
	}
	x := inf
	var sum int
	for i := 0; i < n-1; i++ {
		x = min(x, D[i])
		sum += A[i]
	}
	sum += A[n-1]

	return max(A[n-1]-D[n-1], sum-x, 0)
}

const inf = 1 << 30

func solve1(A []int, D []int) int {
	n := len(A)
	// dp[i] 是前面到i为止的最小D[i]值, 且激活1.。。i为止的结果
	dp := make([]int, n)

	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
		dp[i] = D[i]
		if i > 0 {
			dp[i] = min(dp[i-1], D[i])
		}
	}

	suf := A[n-1]
	sum -= A[n-1]
	var res int
	// best 是后缀种最大的值
	var best int

	for i := n - 2; i > 0; i-- {
		// 激活i+1到结束
		best = max(best, suf-D[i+1])
		// 将i指向0，并激活前半段
		res = max(res, best+sum-dp[i])

		// 将i-1指向i+1, 并激活0和i
		tmp := suf + sum - A[i] - D[0] + max(0, A[i]-D[i])
		res = max(res, tmp)
		// 将E[0] => n-1, 并从i开始
		res = max(res, suf+A[i]-D[i])
		suf += A[i]
		sum -= A[i]
	}

	res = max(res, best)

	return res
}
