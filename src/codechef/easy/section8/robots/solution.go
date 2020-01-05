package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		_, Q := readTwoNums(scanner)
		scanner.Scan()
		S := scanner.Text()
		queries := make([][]int, Q)
		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 2)
		}

		res := solve(S, Q, queries)

		for _, ans := range res {
			a, b := ans[0], ans[1]
			fmt.Printf("%.7f %.7f\n", a, b)
		}
	}
}

func solve(S string, Q int, queries [][]int) [][]float64 {
	P := math.Pi / 3

	// one := complex(1, 0)
	moves := make([]complex128, 6)

	for i := 0; i < 6; i++ {
		moves[i] = cmplx.Rect(1, P*float64(i))
	}

	// cur := complex(0, 0)

	n := len(S)

	dp := make([]complex128, n+1)

	dp[0] = 0 + 0i

	for i := 0; i < n; i++ {
		x := int(S[i] - '0')
		dp[i+1] = dp[i] + moves[x]
	}

	res := make([][]float64, Q)

	for i := 0; i < Q; i++ {
		query := queries[i]
		L := query[0]
		R := query[1]
		x := dp[R] - dp[L-1]
		a := real(x)
		b := imag(x)
		res[i] = []float64{a, b}
	}

	return res
}
