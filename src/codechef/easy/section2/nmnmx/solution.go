package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		A := readNNums(scanner, N)
		res := solve(N, K, A)
		fmt.Printf("%d\n", res)
	}
}

const MOD = 1000000007

const MAX_N = 5000

var C [][]int

func init() {
	// a ^^ x MOD p = a ^^ (x % (p - 1)) MOD P
	C = make([][]int, MAX_N+1)
	C[0] = make([]int, MAX_N+1)
	C[0][0] = 1
	for i := 1; i <= MAX_N; i++ {
		C[i] = make([]int, MAX_N+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % (MOD - 1)
		}
	}
}

func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	return C[n][r]
}

func solve(N, K int, A []int) int {
	X := nCr(N-1, K-1)
	calculate := func(i int) int {
		// A[i] can be used with K - (No of groups that A[i] is min) - (No of groups that A[i] is max)
		// X is the count of groups that has A[i]
		a := nCr(i, K-1)

		b := nCr(N-i-1, K-1)

		return (X - a - b + 2*(MOD-1)) % (MOD - 1)
	}

	// order doesn't affect the way to choose the subsequence
	sort.Ints(A)
	product := int64(1)
	for i := 0; i < N; i++ {
		cnt := calculate(i)
		tmp := int64(pow(A[i], cnt))
		product = (product * tmp) % MOD
	}

	return int(product)
}

func pow(a, b int) int {
	var x = int64(1)
	var y = int64(a)
	for b > 0 {
		if b&1 == 1 {
			x = (x * y) % MOD
		}
		y = (y * y) % MOD
		b >>= 1
	}

	return int(x % MOD)
}
