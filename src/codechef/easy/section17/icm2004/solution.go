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
		n := readNum(scanner)
		scanner.Scan()
		S := scanner.Text()
		solver := NewSolver(n, S)

		m := readNum(scanner)
		for m > 0 {
			m--
			scanner.Scan()
			bs := scanner.Bytes()

			var t, l, r int

			pos := readInt(bs, 0, &t) + 1
			pos = readInt(bs, pos, &l) + 1
			pos = readInt(bs, pos, &r) + 1
			if t == 1 {
				solver.UpdateRange(l, r, bs[pos])
			} else {
				var p, q int
				pos = readInt(bs, pos, &p) + 1
				readInt(bs, pos, &q)
				a, b := solver.Query(l, r, p, q)
				fmt.Printf("%d %d\n", a, b)
			}
		}
	}
}

const MOD = 1000000007

var A, B, I Mat
var PA []Mat
var PB []Mat

const MAX_N = 100007

func init() {
	a := [][]int64{
		{1, 1},
		{MOD - 1, 1},
	}
	b := [][]int64{
		{1, MOD - 1},
		{1, 1},
	}
	A = Mat{a}
	B = Mat{b}

	I = Mat{
		[][]int64{
			{1, 0},
			{0, 1},
		}}

	PA = make([]Mat, MAX_N)
	PB = make([]Mat, MAX_N)

	PA[0] = I
	PB[0] = I

	for i := 1; i < MAX_N; i++ {
		PA[i] = PA[i-1].Mul(A)
		PB[i] = PB[i-1].Mul(B)
	}
}

type Solver struct {
	seg  []Mat
	lazy []int
	n    int
}

func NewSolver(n int, S string) Solver {
	seg := make([]Mat, 4*n)
	lazy := make([]int, 4*n)

	var build func(cur int, st, end int)

	build = func(cur int, st, end int) {
		if st == end {
			if S[st] == 'A' {
				seg[cur] = A
			} else {
				seg[cur] = B
			}
			return
		}
		mid := (st + end) / 2
		build(cur*2, st, mid)
		build(cur*2+1, mid+1, end)
		seg[cur] = seg[cur*2].Mul(seg[cur*2+1])
	}

	build(1, 0, n-1)

	return Solver{seg, lazy, n}
}

func (solver *Solver) UpdateRange(l, r int, C byte) {
	l--
	r--
	seg := solver.seg
	lazy := solver.lazy

	val := 1
	if C == 'B' {
		val = 2
	}

	var loop func(i int, st, end int)

	loop = func(i int, st, end int) {
		if lazy[i] != 0 {
			if lazy[i] == 1 {
				seg[i] = PA[end-st+1]
			} else {
				seg[i] = PB[end-st+1]
			}
			if st != end {
				lazy[2*i] = lazy[i]
				lazy[2*i+1] = lazy[i]
			}
			lazy[i] = 0
		}
		if st > end || st > r || end < l {
			return
		}

		if l <= st && end <= r {
			if val == 1 {
				seg[i] = PA[end-st+1]
			} else {
				seg[i] = PB[end-st+1]
			}

			if st != end {
				lazy[2*i] = val
				lazy[2*i+1] = val
			}
			return
		}

		mid := (st + end) / 2
		loop(2*i, st, mid)
		loop(2*i+1, mid+1, end)
		seg[i] = seg[i*2].Mul(seg[i*2+1])
	}

	loop(1, 0, solver.n-1)
}

func (solver *Solver) Query(l, r int, p, q int) (int, int) {
	l--
	r--
	seg := solver.seg
	lazy := solver.lazy

	var loop func(i int, st, end int) Mat

	loop = func(i int, st, end int) Mat {
		if st > end || st > r || end < l {
			return I
		}
		if lazy[i] != 0 {
			if lazy[i] == 1 {
				seg[i] = PA[end-st+1]
			} else {
				seg[i] = PB[end-st+1]
			}
			if st != end {
				lazy[2*i] = lazy[i]
				lazy[2*i+1] = lazy[i]
			}
			lazy[i] = 0
		}

		if l <= st && end <= r {
			return seg[i]
		}
		mid := (st + end) / 2
		a := loop(2*i, st, mid)
		b := loop(2*i+1, mid+1, end)
		return a.Mul(b)
	}

	a := loop(1, 0, solver.n-1).arr

	P := int64(p)
	Q := int64(q)
	P, Q = a[0][0]*P+a[1][0]*Q, a[0][1]*P+a[1][1]*Q

	return int(P % MOD), int(Q % MOD)
}

type Mat struct {
	arr [][]int64
}

func createMat() [][]int64 {
	arr := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		arr[i] = make([]int64, 2)
	}
	return arr
}

func (this Mat) Mul(other Mat) Mat {
	res := createMat()

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				res[i][j] = (res[i][j] + this.arr[i][k]*other.arr[k][j]) % MOD
			}
		}
	}
	return Mat{res}
}
