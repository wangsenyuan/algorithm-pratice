package main

import (
	"bufio"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	N, Q := readTwoNums(scanner)

	points := make([][]int, N)
	for i := 0; i < N; i++ {
		points[i] = readNNums(scanner, 2)
	}
	queries := make([][]int, Q)
	for i := 0; i < N; i++ {
		queries[i] = readNNums(scanner, 3)
	}

	res := solve(N, Q, points, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}

}

const MAX_X = 300010

func solve(N, Q int, points [][]int, queries [][]int) []int {
	pts := make(AS, N)
	for i := 0; i < N; i++ {
		pts[i] = &Point{points[i][0], points[i][1]}
	}

	sort.Sort(AS(pts))

	qs := make(AS, Q)

	for i := 0; i < Q; i++ {
		q := queries[i]
		qs[i] = &Query{Point: Point{x: q[0], y: q[1]}, d: q[2], index: i}
	}

	sort.Sort(qs)
	res := make([]int, Q)

	xbit := NewBIT(MAX_X)
	ybit := NewBIT(MAX_X)

	for i, j := 0, 0; i < Q; i++ {
		q := qs[i].(*Query)
		for j < N {
			p := pts[j].(*Point)
			if p.x+p.y > q.x+q.y+q.d {
				break
			}
			xbit.Update(p.x, 1)
			ybit.Update(p.y, 1)
			j++
		}
		res[q.index] = j - xbit.Get(q.x-1) - ybit.Get(q.y-1)
	}

	ybit.Reset()

	sort.Sort(BS(pts))
	sort.Sort(BS(qs))

	for i, j := 0, 0; i < Q; i++ {
		q := qs[i].(*Query)
		for j < N {
			p := pts[j].(*Point)
			if p.x >= q.x {
				break
			}
			ybit.Update(p.y, 1)
			j++
		}
		res[q.index] += ybit.Get(q.y - 1)
	}

	return res
}

type BIT struct {
	arr []int
	n   int
}

func NewBIT(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (this *BIT) Update(pos int, val int) {
	pos++
	for pos <= this.n {
		this.arr[pos] += val
		pos += pos & (-pos)
	}
}

func (this BIT) Get(pos int) int {
	pos++
	var res int
	for pos > 0 {
		res += this.arr[pos]
		pos -= pos & (-pos)
	}
	return res
}

func (this *BIT) Reset() {
	for i := 0; i <= this.n; i++ {
		this.arr[i] = 0
	}
}

type A interface {
	X() int
	Y() int
}

type AS []A

func (as AS) Len() int {
	return len(as)
}

func (as AS) Less(i, j int) bool {
	return as[i].X()+as[i].Y() < as[j].X()+as[j].Y()
}

func (as AS) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

type BS AS

func (bs BS) Len() int {
	return len(bs)
}

func (bs BS) Less(i, j int) bool {
	if bs[i].X() < bs[j].X() {
		return true
	}
	if bs[i].X() == bs[j].X() && bs[i].Y() < bs[j].Y() {
		return true
	}
	return false
}

func (bs BS) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

type Point struct {
	x, y int
}

type Query struct {
	Point
	d     int
	index int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}
