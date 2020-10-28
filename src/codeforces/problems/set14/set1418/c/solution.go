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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int) int {
	if n == 1 {
		return A[0]
	}
	// dp[i][0] = min(dp[i-1][1] + x, dp[i-2][1] + x + y)
	// dp[i][1] = dp[i-1][0]

	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = n
		}
	}

	dp[0][0] = A[0]
	dp[1][0] = A[0] + A[1]
	dp[1][1] = A[0]
	//dp[1][0] = A[0] + A[1]
	for i := 2; i < n; i++ {
		dp[idx(i)][1] = min(dp[idx(i-1)][0], dp[idx(i-2)][0])
		dp[idx(i)][0] = min(dp[idx(i-1)][1]+A[i], dp[idx(i-2)][1]+A[i-1]+A[i])
	}

	return min(dp[idx(n-1)][0], dp[idx(n-1)][1])
}

func idx(i int) int {
	i = i % 3
	i = (i + 3) % 3
	return i
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
