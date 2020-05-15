package main

import (
	"bufio"
	"fmt"
	"math"
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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	P := make([][]int, n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		line := readNNums(reader, 3)
		P[i] = []int{line[0], line[1]}
		A[i] = line[2]
	}
	solver := NewSolver(n, P, A)
	Q := readNum(reader)

	for Q > 0 {
		Q--
		line := readNNums(reader, 5)
		if line[0] == 1 {
			x, y := solver.FindAnswer(line[1], line[2], line[3], line[4])
			fmt.Printf("%d %d\n", x, y)
			continue
		}
		solver.UpdatePos(line[1], line[2], line[3], line[4])
	}
}

type Solver struct {
	arr []Matrix
	A   []int
	P   [][]int
	n   int
}

func NewSolver(n int, pts [][]int, A []int) Solver {
	arr := make([]Matrix, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = NewMatrix()
	}

	solver := Solver{arr, A, pts, n}

	for i := 0; i < n; i++ {
		solver.UpdatePos(i+1, pts[i][0], pts[i][1], A[i])
	}

	return solver
}

func (solver *Solver) UpdatePos(i int, x, y, a int) {
	i--
	solver.P[i] = []int{x, y}
	solver.A[i] = a

	val := NewMatrix()

	al := float64(a) * math.Pi / 180

	val[0][0] = int64(math.Cos(al))
	val[1][0] = -int64(math.Sin(al))
	val[2][0] = int64(x)*(1-int64(math.Cos(al))) + int64(y)*int64(math.Sin(al))
	val[0][1] = int64(math.Sin(al))
	val[1][1] = int64(math.Cos(al))
	val[2][1] = int64(y)*(1-int64(math.Cos(al))) - int64(x)*int64(math.Sin(al))
	val[2][2] = 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val[i][j] = (val[i][j]%MOD + MOD) % MOD
		}
	}

	solver.Update(i, val)
}

func (solver *Solver) Update(pos int, v Matrix) {
	pos += solver.n
	solver.arr[pos] = v

	for pos > 1 {
		pos >>= 1
		solver.arr[pos] = solver.arr[pos<<1].Mul(solver.arr[pos<<1|1])
	}
}

func (solver Solver) FindAnswer(x, y, l, r int) (int, int) {
	res := solver.Query(l-1, r)

	xx := (int64(x)*res[0][0] + int64(y)*res[1][0] + res[2][0]) % MOD
	yy := (int64(x)*res[0][1] + int64(y)*res[1][1] + res[2][1]) % MOD

	return int(xx), int(yy)
}

func (solver Solver) Query(l, r int) Matrix {
	lres := UnitMatrix()
	rres := UnitMatrix()

	l += solver.n
	r += solver.n

	for l < r {
		if l&1 == 1 {
			lres = lres.Mul(solver.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			rres = solver.arr[r].Mul(rres)
		}
		l >>= 1
		r >>= 1
	}
	return lres.Mul(rres)
}

const MOD = 1000000007

const S_MOD = MOD * MOD

type Matrix [3][3]int64

func NewMatrix() Matrix {
	x := [3][3]int64{
		[3]int64{0, 0, 0},
		[3]int64{0, 0, 0},
		[3]int64{0, 0, 0},
	}

	return Matrix(x)
}

func UnitMatrix() Matrix {
	res := NewMatrix()
	for i := 0; i < 3; i++ {
		res[i][i] = 1
	}
	return res
}

func (this Matrix) Mul(that Matrix) Matrix {
	res := NewMatrix()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var tmp int64

			for k := 0; k < 3; k++ {
				tmp += this[i][k] * that[k][j]
				if tmp >= S_MOD {
					tmp -= S_MOD
				}
			}

			res[i][j] = tmp % MOD
		}
	}

	return res
}
