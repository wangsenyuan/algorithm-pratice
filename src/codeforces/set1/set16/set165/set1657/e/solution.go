package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	res := solve(n, k)

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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	r %= MOD
	return int(r)
}

const MAX_N = 260

var F [MAX_N]int
var I [MAX_N]int

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r {
		return 0
	}
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(n int, k int) int {
	n--
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1

	for i := 0; i < k; i++ {
		for t := 0; t <= n; t++ {
			pw := pow(k-i, t*(t-1)/2)
			step := pow(k-i, t)
			for j := 0; j < n-t+1; j++ {
				//dp[i][j] = 有j个节点连接到了1，且最大的边的weight是i
				// t个节点使用weight i+1
				// 它们之间可以相互使用(k - (i + 1) + 1)的边进行连接，
				// 这t个节点需要连接到已有的j个节点上，且可以使用(k - i)个weight,
				/**
				 * dp[i+1][j+t] = add(., dp[i][j] * C(n - j, t) * pow(k - i, t * (t - 1) / 2) * pow(k - i, j))
				 **/
				dp[i+1][j+t] = modAdd(dp[i+1][j+t], modMul(dp[i][j], modMul(nCr(n-j, t), pw)))
				pw = modMul(pw, step)
			}
		}
	}

	return dp[k][n]
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
