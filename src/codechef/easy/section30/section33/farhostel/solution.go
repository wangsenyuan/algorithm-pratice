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
		m, n := readTwoNums(reader)
		A := make([][]int, m)
		for i := 0; i < m; i++ {
			A[i] = readNNums(reader, n)
		}
		B := make([][]int, m)
		for i := 0; i < m; i++ {
			B[i] = readNNums(reader, n)
		}
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(A [][]int, B [][]int) int64 {
	m := len(A)
	n := len(A[0])

	// 如果(i, j) go up, 那么在(i, j) 右边的，全部go up
	// 如果(i, j) go leaft, 那么 (i, j)下方的全部 go left
	L := make([][]int64, m+2)
	U := make([][]int64, m+2)

	for i := 0; i < m+2; i++ {
		L[i] = make([]int64, n+2)
		U[i] = make([]int64, n+2)
	}

	// 假设有一个人从(0, 0) 开始，向下，或者向右移动到底部, 这条边界就是向上、或者向左分分界线
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			L[i][j] = int64(A[i-1][j-1]) + L[i][j-1]
			U[i][j] = int64(B[i-1][j-1]) + U[i-1][j]
		}
	}

	dp := make([][]int64, m+2)
	for i := 0; i < m+2; i++ {
		dp[i] = make([]int64, n+2)
	}

	for i := 0; i < m+2; i++ {
		for j := 0; j < n+2; j++ {
			// (i, j) is go up
			if i == 0 || j == 0 {
				continue
			}
			dp[i][j] = max(dp[i-1][j]+L[i][j], dp[i][j-1]+U[i][j])
		}
	}

	return dp[m+1][n+1]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
