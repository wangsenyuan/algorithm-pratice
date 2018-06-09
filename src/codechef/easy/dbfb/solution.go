package main

import (
	"bufio"
	"os"
	"fmt"
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

	t := readNum(scanner)

	for t > 0 {
		M, N := readTwoNums(scanner)
		A := readNNums(scanner, M)
		B := readNNums(scanner, M)
		fmt.Println(solve(M, A, B, N))
		t--
	}
}

const MOD = 1000000007

func solve(M int, A []int, B []int, N int) int64 {
	sa := sum(A)
	sa = (sa * int64(M)) % MOD
	sb := sum(B)
	sb = (sb * int64(M)) % MOD

	if N == 1 {
		return sa
	}

	if N == 2 {
		return sb
	}

	var sc int64
	for i := 3; i <= N; i++ {
		sc = (sa + sb) % MOD
		sa, sb = sb, sc
	}
	return sb
}

func sum(A []int) int64 {
	var res int64

	for i := 0; i < len(A); i++ {
		res = (res + int64(A[i])) % MOD
	}
	return res
}
