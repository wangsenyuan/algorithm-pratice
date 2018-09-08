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

	N, K := readTwoNums(scanner)

	A := readNNums(scanner, N)

	fmt.Println(solve(N, K, A))
}

const LIM = 100111

var I []int

func init() {
	I = make([]int, LIM+1)
}

func solve(N int, K int, A []int) int64 {
	for i := 0; i <= LIM; i++ {
		I[i] = -1
	}
	total := int64(N) * int64(N+1) >> 1
	pi, pj := 0, 1

	addTotal := func(ci, cj int) {
		total += int64(pi) * int64(pj-cj)
		pi = ci
		pj = cj
	}

	var l int
	for j := 1; j <= N; j++ {
		a := A[j-1]
		if a == K {
			l = j
		} else if a > K {
			i := max(I[a], l)
			if pi < i {
				addTotal(i, j)
			}

			a -= K
			for v := 1; v*v <= a; v++ {
				if a%v == 0 {
					I[v] = j
					I[a/v] = j
				}
			}
		}
	}

	addTotal(N, N+1)

	return total
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
