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
		n, x := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, x)
		s := fmt.Sprintf("%d", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

const N_INF = -(1 << 60)

func solve(A []int, x int) []int64 {
	// x >= 0
	// 对于sub array，肯定是把x都加到选中的部分，更好
	//对于给定的k，x可以被添加到k个元素
	// 假设以i结尾的的sub array
	// dp[i][j] = 前i个元素使用了j次后的最大值
	// res[j] = max(res[j], dp[i][j])
	// dp[i+1][j] = max(dp[i][j] + A[i], dp[i][j-1] + A[i] + x)
	n := len(A)

	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
	}

	res := make([]int64, n+1)
	res[0] = 0
	for j := 1; j <= n; j++ {
		res[j] = N_INF
	}
	var sum int64
	var best int64
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			dp[i][j] = max(dp[i-1][j]+int64(A[i-1]), dp[i-1][j-1]+int64(A[i-1])+int64(x))
			res[j] = max(res[j], dp[i][j])
		}

		sum += int64(A[i-1])
		best = max(best, sum)
		if sum < 0 {
			sum = 0
		}
		res[0] = max(res[0], best)
		for j := 1; j <= i; j++ {
			res[j] = max(res[j], res[j-1])
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
