package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		scanner.Scan()
		S := scanner.Text()
		fmt.Println(solve(N, K, S))
	}
}

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

const MAX_N = 100000

var A []int

func init() {
	A = make([]int, MAX_N)
}

func solve(N, K int, S string) int {
	for i := 1; i < N; i++ {
		if S[i] != S[i-1] {
			A[i] = 1
		} else {
			A[i] = 0
		}
	}
	L := N - K
	var X int
	for i := 1; i < L+1; i++ {
		X += A[i]
	}
	ans := X
	for i, j := L+1, 1; i < N; i, j = i+1, j+1 {
		X -= A[j]
		X += A[i]
		ans += X
	}
	return ans
}
