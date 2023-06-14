package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	a := readString(reader)[:n]
	b := readString(reader)[:m]

	res := solve(a, b)

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

func solve(A, B string) int {
	// score(C, D) = 4 * lcs(C, D) - len(C) - len(D)
	// dp[i][j] = (A[:i], B[:j])的最大值
	// dp[i+1][j+1] =
	// lcs(C, D) <= len(C)  and len(D)
	// 2 * lcs(C, D) >=  max(len(C), len(D))
	// 假设 C的长度 >= D
	//  0 < 4 * lcs(C, D) - len(C) - len(D) <= 4 * lcs(C, D) - 2 * len(C)
	// 2 * lcs(C, D) >= len(C)
	// lcs(至少要占一半的C）
	// dp[i, j] 表示前 A[:i], B[:j]的最长lcs
	// dp[i+1, j+1] = dp[i, j] + 1 ? if A[i] == B[j]
	//    else = dp[i, j+1] max dp[i+1 j]
	// score_x(i, j) = 4 * dp[i, j] - i - j
	// score(i, j) = score_x(i, j) - min(score_x(u, v)) where u < i && v < j
	// 计算 score on (a, b) (c, d）
	// 4 * lcs(A[a..b], B[c...d]) - (b - a + 1) - (d - c + 1)
	// 4 * lcs(x, y) - (b + d) - (a + c)
	// lcs(x, y) 是否等于 lcs(A[0...b], B[0...d]) - lca(A[...a], B[...c]) ?
	//   abc
	//   ac

	n := len(A)
	m := len(B)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	var res int

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 2
			} else {
				dp[i][j] = max(0, max(dp[i-1][j], dp[i][j-1])-1)
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
