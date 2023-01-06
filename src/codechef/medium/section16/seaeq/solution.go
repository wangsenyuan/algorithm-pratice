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
	// tc := 1

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
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

const MOD = 1e9 + 7

const N = 510

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
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

var F [N]int
var I [N]int

var dp [N][N * (N - 1) / 2]int
var sum_dp [N][N * (N - 1) / 2]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}
	I[N-1] = pow(F[N-1], MOD-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}

	dp[0][0] = 1
	for i := 1; i < N; i++ {
		m := i * (i - 1) / 2
		var sum int
		for j := 0; j <= m; j++ {
			sum = modAdd(sum, dp[i-1][j])
			dp[i][j] = sum
			sum_dp[i][j] = sum
			if j > 0 {
				sum_dp[i][j] = modAdd(sum_dp[i][j], sum_dp[i][j-1])
			}
			if j >= i {
				sum = modSub(sum, dp[i-1][j-i])
			}
		}
	}
}

func nCr(n int, r int) int {
	if n < r {
		return 0
	}
	res := modMul(F[n], I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(n int, E int) int {
	// every i (1 <= i <= n) CA(A[i]) = CB(B[i]).
	// CX[x] equal to number of index j (1 <=j <= n) such that X[j] < x.
	// A and B 排序是一支的
	// [1, 3, 2] vs [4, 6, 5]
	// P1[l..r] inversion pairs = E
	// P2[l..r] inversion pairs = E
	// 是否等价呢？    1， 3， 2， =1      5, 4, 6 == 1,
	// 不等价
	// F(P1, P2) no of pairs [l..r]
	// 迭代P1[r], 且使得 some l, P1[l..r] = E
	// P2[l..r] 的序列基本可以确定   然后选择 nCr(n, r - l + 1) 按照P1的序列安排，
	// 剩余的部分 （n - (r - l + 1))!
	// l 有哪些呢？ 只要 inversion(r - l + 1) >= E
	// max_inversion(n) = (n - 1) + (n - 2) + .. + 1 = (1 + n - 1) * (n - 1) / 2
	// dp[r][x]
	// E <= n * (n -1) /2

	// dp[i][j]

	var ans int

	for l := 1; l <= n; l++ {
		var tmp int
		if E > l*(l-1)/2 {
			tmp = sum_dp[l][l*(l-1)/2]
		} else {
			tmp = sum_dp[l][E]
		}

		tmp = modMul(tmp, F[n-l])
		tmp = modMul(tmp, F[n-l])
		tmp = modMul(tmp, n-l+1)
		tmp = modMul(tmp, nCr(n, l))
		tmp = modMul(tmp, nCr(n, l))

		ans = modAdd(ans, tmp)
	}

	return ans
}
