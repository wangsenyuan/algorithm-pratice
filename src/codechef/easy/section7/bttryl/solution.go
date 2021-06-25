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
		n := readNum(reader)
		res := solve(n)
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

const N = 110
const MOD = 1000000007

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}

	return int(R)
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

var I []int64

func init() {
	I = make([]int64, N)

	for i := 1; i < N; i++ {
		I[i] = int64(inverse(i))
	}
}

func solve(n int) int {
	dp := make([][][]int64, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int64, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var F func(n, p, q int) int64

	F = func(n, p, q int) int64 {
		if n == 1 || n-p-q < 0 {
			return 0
		}
		if dp[n][p][q] >= 0 {
			return dp[n][p][q]
		}
		dp[n][p][q] = 0
		if q == 0 && p > 0 {
			dp[n][p][q] = modAdd(1, F(p, 0, p))
		} else {
			// first in Q, kill someone still alive in P, then p decrease one, but add the first
			var res int64
			if p > 0 {
				res = modMul(modMul(int64(p), I[n-1]), F(n, p, q-1))
			}
			// first in Q, kill someone still alive in Q, then q decrease two
			if q >= 2 {
				res = modAdd(res, modMul(modMul(int64(q-1), I[n-1]), F(n, p+1, q-2)))
			}
			// or first in Q, kill someone already out
			res = modAdd(res, modMul(modMul(int64(n-q-p), I[n-1]), F(n, p+1, q-1)))
			dp[n][p][q] = res
		}

		return dp[n][p][q]
	}

	return int(F(n, 0, n))
}

func modMul(a, b int64) int64 {
	a *= b
	return a % MOD
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
