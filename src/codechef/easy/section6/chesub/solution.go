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
		N, K := readTwoNums(reader)
		A := readNNums(reader, N)

		res := solve(N, K, A)

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

const INF = int64(1) << 50
const N_INF = -INF

func solve(N, K int, A []int) int64 {
	dp := make([][]int64, N)
	pf := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, K+1)
		pf[i] = make([]int64, K+1)
		for j := 1; j <= K; j++ {
			dp[i][j] = N_INF
			pf[i][j] = N_INF
		}

	}

	// dp[i][j] = max(dp[i-1][j] + j * A[i], pf[i-1][j-1] + j * A[i])
	// pf[i][j] = max(pf[i-1][j], dp[i][j])
	dp[0][1] = int64(A[0])
	pf[0][1] = dp[0][1]
	for i := 1; i < N; i++ {
		for j := 1; j <= K; j++ {
			dp[i][j] = max(dp[i-1][j], pf[i-1][j-1]) + int64(j)*int64(A[i])
			pf[i][j] = max(pf[i-1][j], dp[i][j])
		}
	}

	var res = N_INF

	for i := 0; i < N; i++ {
		res = max(res, dp[i][K])
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
