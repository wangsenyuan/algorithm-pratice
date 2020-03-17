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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, q := readTwoNums(scanner)
		fillNNums(scanner, n, A)

		solver := NewSolver(n, A)

		for q > 0 {
			q--
			p := readNum(scanner)
			fmt.Println(solver.Ask(p))
		}
	}
}

const MAX_N = 100001

var A []int

func init() {
	A = make([]int, MAX_N)
}

type Solver struct {
	even, odd int
}

func NewSolver(n int, A []int) Solver {
	var odd, even int

	for i := 0; i < n; i++ {
		var cnt int
		for A[i] > 0 {
			cnt += A[i] & 1
			A[i] >>= 1
		}
		if cnt&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return Solver{odd, even}
}

func (solver Solver) Ask(p int) (x int, y int) {
	var cnt int
	for p > 0 {
		cnt += p & 1
		p >>= 1
	}
	// o + o => e
	// o + e => o
	// e + e => e
	// e + o => o
	if cnt&1 == 1 {
		// odd
		x = solver.even
		y = solver.odd
	} else {
		x = solver.odd
		y = solver.even
	}
	return
}
