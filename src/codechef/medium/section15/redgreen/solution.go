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
		n, m := readTwoNums(reader)
		res := solve(n, m)
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

const MOD = 998244353

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

const MAX_N = 2010

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
	if n < r || r < 0 {
		return 0
	}
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(n int, m int) int {
	// total 2 ** (n * m) ways
	// n - 1, m 且最后一个格子是
	// dp[i][j][x] = 表示到[i, j]， 且path里面有x个红色时的路径的计数
	// dp[i+1][j][x] = dp[i][j][x] (i+1, j) + dp[i+1][j-1][x] 标黄色
	//               + dp[i][j][x - 1] (i + 1, j) + dp[i+1][j-1][x-1] 标红色
	if (n+m-1)%2 != 0 {
		return 0
	}
	// C(n + m - 1, n) * C(n, p) where p = (n + m - 1) / 2

	// 选择n个down
	x := nCr(n-1+m-1, n-1)

	y := nCr(n+m-1, (n+m-1)/2)

	z := pow(2, n*m-(n+m-1))

	res := modMul(x, y)
	res = modMul(res, z)

	return res
}
