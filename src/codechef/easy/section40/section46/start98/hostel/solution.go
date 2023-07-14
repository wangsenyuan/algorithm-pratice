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
		n, m, x := readThreeNums(reader)
		A := readNNums(reader, n)
		H := make([][]int, n)
		for i := 0; i < n; i++ {
			H[i] = readNNums(reader, m)
		}
		res := solve(m, x, A, H)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

const inf = 1 << 62

func solve(m int, x int, A []int, H [][]int) int64 {
	n := len(A)
	// 如果对于student i 和 student j
	// 现在i所在的房间是 a[i], a[j]
	// 如果i顺时针移动x距离后，到达 (a[i] + x) % n房间后
	//  H[i][a[i]] + H[j][b[j]] > ... 那么这是一个划算的操作
	// i能到达的位置，是定的  gcd(m, x) = 1, 可以到达所有的位置， 否则的话
	// 假设有w次操作，
	// 那么就肯定是w次正向操作，w次反向操作
	// 假设 dp[w] 是正向操作后的最大值, fp[w]是反向操作后的最大值, dp[w] + fp[w]
	// 可以只考虑一个方向的操作
	// 在w确定的情况下，如何计算最优的结果呢？
	// 动态规划 dp[i][x] 表示前i个元素进行了x次正向操作后的，最大值
	// dp[i+1][x+1] = max(dp[i][x+1] + H[i][A[j]], dp[i][x] + x)
	// m次后，肯定回到原地
	ln := m / gcd(m, x)

	dp := make([]int64, ln)
	for i := 0; i < ln; i++ {
		dp[i] = -inf
	}

	dp[0] = 0

	fp := make([]int64, ln)

	for i := 0; i < n; i++ {
		pos := A[i] - 1
		for j := 0; j < ln; j++ {
			fp[j] = -inf
		}

		for j := 0; j < ln; j++ {
			for k := 0; k < ln; k++ {
				where := (pos + k*x) % m
				fp[(j+k)%ln] = max(fp[(j+k)%ln], dp[j]+int64(H[i][where]))
			}
		}
		copy(dp, fp)
	}

	return dp[0]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
