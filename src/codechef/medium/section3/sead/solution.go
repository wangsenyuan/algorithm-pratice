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

	n := readNum(scanner)

	A := readNNums(scanner, n)

	m := readNum(scanner)

	quries := make([][]int, m)
	for i := 0; i < m; i++ {
		quries[i] = readNNums(scanner, 2)
	}

	res := solve(n, A, m, quries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

var MAX_N = 100000
var LOG []int

func init() {
	LOG = make([]int, MAX_N+1)
	LOG[1] = 0
	for i := 2; i <= MAX_N; i++ {
		LOG[i] = LOG[i/2] + 1
	}
}

func solve(n int, A []int, m int, quries [][]int) []int {

	B := make([]int, n-1)

	P := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		P[i] = make([]int, 15)
	}

	for i := 0; i < n-1; i++ {
		B[i] = A[i+1] - A[i]
		P[i][0] = B[i]
	}

	for j := 1; j < 15; j++ {
		for i := 0; i < n-1; i++ {
			k := i + 1<<uint(j-1)
			if k < n-1 {
				P[i][j] = max(P[i][j-1], P[k][j-1])
			}
		}
	}

	G := func(L, R, D int) bool {
		// max value in B[L...R] <= D
		lg := LOG[R-L]
		x := P[L][lg]
		y := P[R-1<<uint(lg)][lg]
		x = max(x, y)

		return x <= D
		// f := math.Log2(float64(R - L))
		// i := int(f)

		// return P[L][i] <= D
	}

	res := make([]int, m)

	for i := 0; i < m; i++ {
		t, d := quries[i][0], quries[i][1]
		R := sort.Search(n, func(j int) bool {
			return A[j] > t
		})
		R--
		a, b := 0, R

		for a < b {
			mid := (a + b) >> 1
			if G(mid, R, d) {
				b = mid
			} else {
				a = mid + 1
			}
		}
		res[i] = a + 1
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
