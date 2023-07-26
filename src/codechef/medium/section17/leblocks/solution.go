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
		n, k := readTwoNums(reader)
		A := make([]int, n)
		C := make([]int, n)
		for i := 0; i < n; i++ {
			A[i], C[i] = readTwoNums(reader)
		}
		res := solve(A, C, k)
		buf.WriteString(fmt.Sprintf("%.8f\n", res))
	}

	fmt.Print(buf.String())
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MAX_N = 30

var F [MAX_N]int64

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1]
	}
}

func solve(A []int, C []int, k int) float64 {
	n := len(A)
	// n <= 20
	// A[i] <= 10
	// C[i] <= n
	// k <= 200
	// if k > sum(A) => 0
	// i, j, 如果 C[i] = C[j]，那么在它们中间如果能放置长度为 k - A[i]的（不包括i, j)的方案数
	// 剩余的其他的放置在i, j的两端
	count := func(a, b, d int) int {
		l := a + d
		r := a + b + d - 1
		ll := k
		rr := k + a - 1
		return max(0, min(r, rr)-max(l, ll)+1)
	}

	var tot int
	for _, a := range A {
		tot += a
	}

	vis := make([]bool, n)
	var ans float64

	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, tot+1)
	}

	process := func(a int, b int) {
		for i := 0; i <= n; i++ {
			for j := 0; j <= tot; j++ {
				dp[i][j] = 0
			}
		}
		dp[0][0] = 1
		var sum int
		for j := 0; j < n; j++ {
			if j == a || j == b {
				continue
			}
			for nj := j; nj >= 0; nj-- {
				for s := sum; s >= 0; s-- {
					dp[nj+1][s+A[j]] += dp[nj][s]
				}
			}
			sum += A[j]
		}

		for j := 0; j < n-1; j++ {
			for s := 0; s <= sum; s++ {
				p := dp[j][s] * float64(F[j]) * float64(j+1) * float64(F[n-j-2]) * 2
				p /= float64(F[n])
				p *= float64(count(A[a], A[b], tot-A[a]-A[b]-s))
				ans += p
			}
		}
	}

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if C[j] == C[i] {
				vis[j] = true
				ans += float64(max(0, A[j]-k))
				process(i, j)

				break
			}
		}
		ans += float64(max(0, A[i]-k))
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
