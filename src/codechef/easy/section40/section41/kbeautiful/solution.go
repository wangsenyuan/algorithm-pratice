package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		A := readNInt64s(reader, n)
		res := solve(k, m, A)
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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

const MAX_N = 2000010

var F []int
var IF []int

func init() {
	F = make([]int, MAX_N)

	F[0] = 1

	for i := 1; i < MAX_N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	IF = make([]int, MAX_N)
	IF[MAX_N-1] = inverse(F[MAX_N-1])

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = modMul(i+1, IF[i+1])
	}
}

func nCr(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	res := F[n]
	res = modMul(res, IF[r])
	res = modMul(res, IF[n-r])
	return res
}

func solve(K int, M int, A []int64) int {
	// A[k] = A[0], A[k+1] = A[1] ... and so one

	make_equal := func(arr []int64) int64 {
		var x int64
		var sum int64
		for i := 0; i < len(arr); i++ {
			if x < arr[i] {
				x = arr[i]
			}
			sum += arr[i]
		}
		return int64(len(arr))*x - sum
	}

	buf := make([][]int64, K)
	n := len(A)
	for i := 0; i < n; i++ {
		j := i % K
		if j == i {
			buf[j] = make([]int64, 0, (n+K-1)/K)
		}
		buf[j] = append(buf[j], A[i])
	}
	var sum int64
	for i := 0; i < K; i++ {
		sum += make_equal(buf[i])
	}

	if sum > int64(M) {
		return 0
	}
	// have ways to handle, sum <= M
	M -= int(sum)
	// 还有这么多操作可以进行
	// L0 * x0 + L1 * x1 + .. L(k-1) * x(k-1) <= M
	// 在多维空间中如何计算呢?
	// x[i]是第i个分组进行的操作次数
	// x[i] <= X[i], 且 sum(L[i] * x[i]) <= M
	// dp[i][m]动态规划是可以的，但是 k * m 太大了
	// 这里有个关键的信息 L[i]最多相差1
	ct := K - n%K
	lo := n / K
	hi := (n + K - 1) / K

	f := func(x, k int) int {
		if x == 0 {
			if k == 0 {
				return 1
			}
			return 0
		}

		return nCr(x+k-1, k)
	}

	dp := make([]int, M+1)

	dp[0] = f(K-ct, 0)

	for i := 1; i <= M; i++ {
		dp[i] = modAdd(dp[i-1], f(K-ct, i))
	}

	var ans int
	for i := 0; i <= M/lo; i++ {
		y := M - lo*i
		ans = modAdd(ans, modMul(f(ct, i), dp[y/hi]))
	}

	return ans
}
