package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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

	r := readNum(scanner)

	for r > 0 {
		tmp := readNNums(scanner, 4)
		fmt.Println(solve(tmp[0], tmp[1], tmp[2], tmp[3]))
		r--
	}
}

const MAX_N = 800001
const MOD = 1000000007

var fact []int64
var ifact []int64

func init() {
	inv := make([]int64, MAX_N)
	inv[1] = 1
	for i := 2; i < MAX_N; i++ {
		inv[i] = MOD - ((MOD / int64(i)) * inv[MOD%i] % MOD)
	}

	fact = make([]int64, MAX_N)
	ifact = make([]int64, MAX_N)

	fact[0] = 1
	ifact[0] = 1

	for i := 1; i < MAX_N; i++ {
		fact[i] = (int64(i) * fact[i-1]) % MOD
		// ifact[i] = pow(fact[i], MOD-2) % MOD
		ifact[i] = (ifact[i-1] * inv[i]) % MOD
	}
}

func pow(a, b int64) int64 {
	if b == 0 {
		return 1
	}

	c := pow(a, b/2)
	d := (c * c) % MOD
	if b%2 == 1 {
		return (a * d) % MOD
	}
	return d
}

func C(n, m int) int64 {
	return ((fact[n] * ifact[n-m]) % MOD * ifact[m]) % MOD
}

func ways(x, y int) int64 {
	return C(x+y, y)
}

func solve(N, M, A, B int) int64 {
	if A == 0 || B == 0 {
		return ways(N, M)
	}
	var res int64
	for p := 0; p <= N-A; p++ {
		res = (res + ways(p, B-1)*ways(N-p, M-B)) % MOD
	}
	return res
}
