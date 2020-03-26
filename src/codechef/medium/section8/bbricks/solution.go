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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		fmt.Println(solve(readTwoNums(scanner)))
	}
}

const MOD = 1000000007
const MAX_K = 1007

var DP [][]int64

var F []int64
var IF []int64

func init() {
	F = make([]int64, MAX_K)
	F[0] = 1
	for i := 1; i < MAX_K; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}

	IF = make([]int64, MAX_K)
	IF[MAX_K-1] = pow(F[MAX_K-1], MOD-2)
	for i := MAX_K - 2; i >= 0; i-- {
		IF[i] = (int64(i+1) * IF[i+1]) % MOD
	}

	DP = make([][]int64, MAX_K)
	for i := 0; i < MAX_K; i++ {
		DP[i] = make([]int64, MAX_K)
	}
	DP[1][0] = 2
	for i := 2; i < MAX_K; i++ {
		DP[i][0] = 2
		for j := 1; j < MAX_K; j++ {
			DP[i][j] = DP[i-1][j] + DP[i-1][j-1]
			DP[i][j] %= MOD
		}
	}
}

func pow(a int64, b int) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
			res %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return res
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[r]
	res %= MOD
	res *= IF[n-r]
	res %= MOD
	return res
}

func nCr2(n int, r int) int64 {
	if n < MAX_K {
		return nCr(n, r)
	}
	low := n - r
	var cur int64 = 1
	for n > low {
		cur *= int64(n)
		cur %= MOD
		n--
	}
	cur *= IF[r]
	cur %= MOD
	return cur
}

func solve(n, k int) int {
	var res int64

	for i := 0; i < k; i++ {
		cur := DP[k][i] * nCr2(n-i, k)
		cur %= MOD
		res += cur
		res %= MOD
	}

	return int(res)
}
