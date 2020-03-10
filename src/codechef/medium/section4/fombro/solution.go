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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		n, m, q := line[0], line[1], line[2]
		solver := NewSolver(n, m)

		for i := 0; i < q; i++ {
			r := readNum(scanner)
			fmt.Println(solver.Query(r))
		}
	}
}

type Solver struct {
	fordp, beddp, backdp []int64
	n                    int
	m                    int
}

const MAX_N = 500002

var fordp []int64
var beddp []int64
var backdp []int64

func init() {
	fordp = make([]int64, MAX_N)
	beddp = make([]int64, MAX_N)
	backdp = make([]int64, MAX_N)
}

func pow(a, b, m int) int64 {
	A := int64(a)
	r := int64(1)
	M := int64(m)

	for b > 0 {
		if b&1 == 1 {
			r *= A
			r %= M
		}
		A = (A * A) % M
		b >>= 1
	}
	return r
}

func NewSolver(n, m int) Solver {
	fordp[0] = 1
	backdp[0] = 1

	M := int64(m)

	for i := 1; i <= n/2; i++ {
		fordp[i] = (fordp[i-1] * pow(i, i-1, m)) % M
		backdp[i] = (backdp[i-1] * pow(n-i+1, i, m)) % M
	}

	if n&1 == 1 {
		beddp[n/2] = int64(n/2+1) % M
	} else {
		beddp[n/2] = 1
	}

	for i := n/2 - 1; i >= 1; i-- {
		beddp[i] = (beddp[i+1] * int64(i+1)) % M
		beddp[i] *= int64(n - i)
		beddp[i] %= M
	}

	return Solver{fordp, beddp, backdp, n, m}
}

func (solver Solver) Query(r int) int64 {
	r = min(r, solver.n-r)

	ans := solver.fordp[r]
	ans *= pow(int(solver.beddp[r]), r, solver.m)
	ans %= int64(solver.m)
	ans *= solver.backdp[r]
	ans %= int64(solver.m)

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
