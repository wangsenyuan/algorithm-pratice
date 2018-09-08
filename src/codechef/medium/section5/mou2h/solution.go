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
	res := A
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res[:n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N := readNum(scanner)
		X := readNNums(scanner, N)
		res := solve(N, X)
		fmt.Println(res)
	}
}

const MOD = 1e9 + 9
const MAX_H = 4 * 1e6
const MAX_N = 1e6 + 10

var H []int
var A []int
var B []int
var C []int
var dp []int64

func init() {
	H = make([]int, MAX_H*2+10)
	A = make([]int, MAX_N)
	B = make([]int, MAX_N)
	C = make([]int, MAX_N)
	dp = make([]int64, MAX_N)
}

func solve(N int, A []int) int64 {
	for i := 0; i < MAX_H; i++ {
		H[i] = -1
	}
	var idx int
	for i := 1; i < N; i++ {
		B[i-1] = A[i] - A[i-1]
		if B[i-1] < 0 {
			B[i-1] += MAX_H
		}
		if H[B[i-1]] < 0 {
			H[B[i-1]] = idx
			idx++
		}
	}

	for i := 0; i < N-1; i++ {
		B[i] = H[B[i]]
	}

	for i := 0; i < idx; i++ {
		C[i] = -1
	}

	dp[0] = 1

	for i := 0; i < N-1; i++ {
		dp[i+1] = (2 * dp[i]) % MOD
		if C[B[i]] >= 0 {
			dp[i+1] -= dp[C[B[i]]]
			if dp[i+1] < 0 {
				dp[i+1] += MOD
			}
		}
		C[B[i]] = i
	}
	res := dp[N-1] - 1
	if res < 0 {
		res += MOD
	}
	return res
}
