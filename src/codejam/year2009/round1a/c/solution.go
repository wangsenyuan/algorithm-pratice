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

	for x := 1; x <= tc; x++ {
		C, N := readTwoNums(scanner)
		res := solve(C, N)
		fmt.Printf("Case #%d: %.6f\n", x, res)
	}
}

const MAX_N = 40

var X [][]int64

func init() {
	X = make([][]int64, MAX_N+1)

	for i := 0; i <= MAX_N; i++ {
		X[i] = make([]int64, MAX_N+1)
		X[i][0] = 1
		X[i][i] = 1
		for j := 1; j < i; j++ {
			X[i][j] = X[i-1][j] + X[i-1][j-1]
		}
	}
}

func nOc(a, b int) int64 {
	if b < 0 || b > a {
		return 0
	}
	return X[a][b]
}

func solve(C, N int) float64 {
	c := nOc(C, N)

	E := make([]float64, C+1)
	E[0] = 0

	for i := 1; i <= C; i++ {
		var sum float64
		for j := 1; j <= N && j <= i; j++ {
			tmp := float64(X[i][j]*X[C-i][N-j]) / float64(c)
			sum += tmp * (1.0 + E[i-j])
		}
		zero := float64(X[C-i][N]) / float64(c)
		E[i] = (sum + zero) / (1.0 - zero)
	}

	return E[C]
}
