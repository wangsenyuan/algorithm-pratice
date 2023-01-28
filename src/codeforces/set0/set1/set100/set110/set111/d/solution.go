package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	res := solve(n, m, k)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
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

const MAX_K = 1e6 + 10

var F [MAX_K]int
var I [MAX_K]int

func init() {
	F[0] = 1
	for i := 1; i < MAX_K; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I[MAX_K-1] = pow(F[MAX_K-1], MOD-2)
	for i := int(MAX_K - 2); i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	return modMul(modMul(F[n], I[r]), I[n-r])
}

func solve(n int, m int, k int) int {
	/* for any vertical line that passes along the grid
	lines and divides the board in two non-empty parts the number
	of distinct colors in both these parts should be the same. */
	if m == 1 {
		return pow(k, n)
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	// d[i][j] = j * d[i — 1][j] + d[i — 1][j — 1].
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = modAdd(modMul(j, dp[i-1][j]), dp[i-1][j-1])
		}
	}
	var ans int
	// y 是第一列和最后一列相同的颜色数
	for y := 0; y <= n; y++ {
		cur := pow(y, n*(m-2))
		// x 是分割后，两边（不同颜色）的树目
		for x := y; x <= n; x++ {
			if 2*x-y <= k {
				tek := nCr(k, 2*x-y)
				// 第一列从 2 * x - y 中选择 x- y
				tek = modMul(tek, nCr(2*x-y, x-y))
				// 最后一列从x（不同颜色的树木）选择 (x - y) 只在最后一列出现的颜色
				tek = modMul(tek, nCr(x, x-y))
				tek = modMul(tek, pow(dp[n][x], 2))
				tek = modMul(tek, pow(F[x], 2))
				tek = modMul(tek, cur)
				ans = modAdd(ans, tek)
			}
		}
	}

	return ans
}
