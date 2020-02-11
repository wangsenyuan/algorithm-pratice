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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		solver := NewSolver(n, A)

		for q > 0 {
			q--
			m := readNum(scanner)
			fmt.Println(solver.Query(m))
		}

	}
}

const MOD = 1000000009

const MAX_N = 100005

var fact []int64
var ifact []int64

func init() {
	fact = make([]int64, MAX_N)

	fact[0] = 1
	for i := 1; i < MAX_N; i++ {
		fact[i] = (int64(i) * fact[i-1]) % MOD
	}

	ifact = make([]int64, MAX_N)
	ifact[MAX_N-1] = pow(int(fact[MAX_N-1]), MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		ifact[i] = int64(i+1) * ifact[i+1]
		ifact[i] %= MOD
	}
}

func pow(a, b int) int64 {
	r := int64(1)
	A := int64(a)

	for b > 0 {
		if b&1 == 1 {
			r *= A
			r %= MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return r
}

type Solver struct {
	A     []int
	dp    [][]int64
	sum   []int64
	count []int
}

const MAX_M = 100

func NewSolver(n int, A []int) Solver {
	dp := make([][]int64, MAX_M+1)

	for i := 0; i <= MAX_M; i++ {
		dp[i] = make([]int64, MAX_M)
	}

	sum := make([]int64, MAX_M)

	count := make([]int, MAX_M)

	return Solver{A, dp, sum, count}
}

func (solver *Solver) clean() {
	for i := 0; i <= MAX_M; i++ {
		for j := 0; j < MAX_M; j++ {
			solver.dp[i][j] = 0
		}
	}

	for i := 0; i < MAX_M; i++ {
		solver.sum[i] = 0
		solver.count[i] = 0
	}

}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := fact[n]
	res *= ifact[r]
	res %= MOD
	res *= ifact[n-r]
	res %= MOD
	return res
}

func (solver *Solver) Query(m int) int {
	solver.clean()

	A := solver.A

	n := len(A)

	count := solver.count

	for i := 0; i < n; i++ {
		a := A[i] % m
		if a < 0 {
			a += m
		}
		count[a]++
	}

	sum := solver.sum
	dp := solver.dp
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			sum[j] = 0
		}
		for k := 0; k <= count[i]; k++ {
			sum[(k*i)%m] += nCr(count[i], k)
			sum[(k*i)%m] %= MOD
		}
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				tmp := dp[i][j] * sum[k]
				tmp %= MOD
				dp[i+1][(j+k)%m] += tmp
				dp[i+1][(j+k)%m] %= MOD
			}
		}

	}

	return int(dp[m][0])
}
