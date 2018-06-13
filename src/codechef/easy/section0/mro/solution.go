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

const MOD = 1e9 + 7
const MAX_N = 1e5

var pw2 []int64
var f []int64

func init() {
	pw2 = make([]int64, MAX_N+1)

	pw2[0] = 1
	for i := 1; i <= MAX_N; i++ {
		pw2[i] = (2 * pw2[i-1]) % MOD
	}
	f = make([]int64, MAX_N+1)
	f[1] = 1
	f[2] = 1
	for i := 3; i <= MAX_N; i++ {
		f[i] = (f[i-1] * (pw2[i-1] - 1 + MOD)) % MOD
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		fmt.Println(solve(n))
		t--
	}
}

func solve(n int) int64 {
	return f[n]
}
