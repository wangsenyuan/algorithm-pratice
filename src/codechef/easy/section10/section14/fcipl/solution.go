package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		p, q, r := line[0], line[1], line[2]
		fmt.Println(solve(p, q, r))
	}
}

const MOD = 998244353

const N = 101

var fact []int64
var ifact []int64
var dp [][]int64
var fp [][]int64

func init() {
	fact = make([]int64, N)
	fact[0] = 1
	for i := 1; i < N; i++ {
		fact[i] = int64(i) * fact[i-1]
		fact[i] %= MOD
	}

	ifact = make([]int64, N)

	ifact[N-1] = pow(fact[N-1], MOD-2)

	for i := N - 2; i >= 0; i-- {
		ifact[i] = int64(i+1) * ifact[i+1]
		ifact[i] %= MOD
	}

	nCr := func(a, b int) int64 {
		if a < b {
			return 0
		}
		res := fact[a]
		res *= ifact[b]

		res %= MOD

		res *= ifact[a-b]

		res %= MOD

		return res
	}

	dp = make([][]int64, N)
	fp = make([][]int64, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int64, N)
		fp[i] = make([]int64, N)
	}

	dp[0][0] = 1
	fp[0][0] = 1

	for n := 1; n < N; n++ {
		for r := 1; r < N; r++ {
			// try dp[n][r] & fp[n][r]
			var tmp int64
			for k := 0; k < n; k++ {
				// fix 1 at last position, try to pick k footballers from previous n - 1 ones to pair with the last one
				tmp += nCr(n-1, k) * fp[n-1-k][r-1]
				tmp %= MOD
			}
			fp[n][r] = tmp

			tmp = 0

			for k := 1; k < n; k++ {
				tmp += nCr(n-1, k) * dp[n-1-k][r-1]
				tmp %= MOD
			}

			dp[n][r] = tmp
		}
	}

}

func pow(A int64, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res *= A
			res %= MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return res
}

func solve(P, Q, R int) int64 {
	// divides Q into r (< R) groups
	// Q! * ?
	// how to make sure at least 2 in one group
	// ** | *** | ** | **

	// divides P into r (< R) groups
	// fact[P] * nCr(P - 1, r - 1)

	var res int64

	for x := 1; x < R; x++ {
		y := R - x
		tmp := fp[P][x]
		tmp %= MOD
		tmp *= dp[Q][y]
		tmp %= MOD
		res += tmp
		res %= MOD
	}

	return res
}
