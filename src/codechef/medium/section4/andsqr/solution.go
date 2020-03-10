package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		n, Q := readTwoNums(scanner)
		A := readNNums(scanner, n)

		queries := make([][]int, Q)

		for i := 0; i < Q; i++ {
			queries[i] = readNNums(scanner, 2)
		}

		res := solve(n, A, Q, queries)

		for i := 0; i < Q; i++ {
			fmt.Println(res[i])
		}
	}
}

var bit2 *BIT2

const MAX_N_1 = 1000100

var qrs []Query

var fsu [][]int

func init() {
	tmp := (NewBIT2(MAX_N_1))
	bit2 = &tmp
	qrs = make([]Query, MAX_N_1)
	fsu = make([][]int, MAX_N_1)
	for i := 0; i < MAX_N_1; i++ {
		fsu[i] = make([]int, 31)
	}
}

func solve(n int, A []int, Q int, queries [][]int) []int64 {
	bit2.Reset(n + 1)

	for i := 0; i < Q; i++ {
		qrs[i] = Query{queries[i][0], queries[i][1], i}
	}

	sort.Sort(Queries(qrs[:Q]))

	for j := 0; j < 31; j++ {
		fsu[n][j] = n + 1
	}

	for i := n - 1; i > 0; i-- {
		for j := 0; j < 31; j++ {
			if (1<<uint(j))&A[i] == 1 {
				fsu[i][j] = fsu[i+1][j]
			} else {
				fsu[i][j] = i + 1
			}
		}
	}

	cl := n + 1
	ans := make([]int64, Q)

	// l is decreasing
	for i := 0; i < Q; i++ {
		l, r := qrs[i].l, qrs[i].r

		for cl > l {
			cl--
			y := A[cl-1]
			z := cl

			for z <= n {
				y &= A[z-1]

				if y == 0 {
					bit2.UpdateRange(n, z-1, n-1, 1)
					break
				}
				nx := n + 1
				for j := 0; j < 31; j++ {
					if y&(1<<uint(j)) > 0 {
						nx = min(nx, fsu[z][j])
					}
				}
				m := int(math.Sqrt(float64(y)))

				if m*m == y {
					bit2.UpdateRange(n, z-1, nx-2, 1)
				}
				z = nx
			}
		}
		ans[qrs[i].id] = bit2.SumRange(n, r-1)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Query struct {
	l, r int
	id   int
}

type Queries []Query

func (this Queries) Len() int {
	return len(this)
}

func (this Queries) Less(i, j int) bool {
	return this[i].l > this[j].l
}

func (this Queries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type BIT struct {
	arr []int64
}

func NewBIT(n int) BIT {
	arr := make([]int64, n+1)
	return BIT{arr}
}

func (bit *BIT) Update(n int, pos int, v int64) {
	pos++

	for pos <= n {
		bit.arr[pos] += v
		pos += pos & -pos
	}
}

func (bit BIT) Get(n int, pos int) int64 {
	var res int64

	pos++

	for pos > 0 {
		res += bit.arr[pos]
		pos -= pos & -pos
	}

	return res
}

func (bit *BIT) Reset(n int) {
	for i := 0; i <= n; i++ {
		bit.arr[i] = 0
	}
}

type BIT2 struct {
	bit1 *BIT
	bit2 *BIT
}

func NewBIT2(n int) BIT2 {
	bit1 := NewBIT(n)
	bit2 := NewBIT(n)

	return BIT2{&bit1, &bit2}
}

func (bit2 BIT2) SumRange(n int, x int) int64 {
	X := int64(x)
	res := bit2.bit1.Get(n, x) * X
	res -= bit2.bit2.Get(n, x)
	return res
}

func (bit2 *BIT2) UpdateRange(n int, l, r int, v int64) {
	bit2.bit1.Update(n, l, v)
	bit2.bit1.Update(n, r+1, -v)

	bit2.bit2.Update(n, l, v*int64(l-1))
	bit2.bit2.Update(n, r+1, -v*int64(r))
}

func (bit2 BIT2) Reset(n int) {
	bit2.bit1.Reset(n)
	bit2.bit2.Reset(n)
}
