package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		solver := NewSolver(n, P)
		for q > 0 {
			q--
			x, y := readTwoNums(reader)
			res := solver.Query(x, y)
			buf.WriteString(fmt.Sprintf("%d\n", res))
		}
	}
	fmt.Print(buf.String())
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

type Solver struct {
	layers []Points
}

func NewSolver(n int, P [][]int) Solver {
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		cur := P[i]
		points[i] = Point{cur[0], cur[1], i}
	}
	used := make(map[int]bool)
	var layers []Points
	for n >= 3 {
		hull := convexHull(points[:n])
		layers = append(layers, hull)
		for i := 0; i < len(hull); i++ {
			used[hull[i].id] = true
		}
		var j int
		for i := 0; i < n; i++ {
			if !used[points[i].id] {
				points[j] = points[i]
				j++
			}
		}
		n = j
	}

	return Solver{layers}
}

func (solver Solver) Query(x, y int) int {
	var left, right = 0, len(solver.layers)
	p := Point{x, y, -1}
	for left < right {
		mid := (left + right) / 2
		if !solver.layers[mid].Contains(p) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

type Point struct {
	x, y int
	id   int
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y, -1}
}

func (this Point) Cross(that Point) int {
	a := int64(this.x) * int64(that.y)
	b := int64(this.y) * int64(that.x)
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func cross(a, b, c Point) int {
	return b.Sub(a).Cross(c.Sub(a))
}

type Points []Point

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	return ps[i].y < ps[j].y || ps[i].y == ps[j].y && ps[i].x < ps[j].x
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Points) Contains(point Point) bool {
	for i := 0; i < len(ps); i++ {
		j := (i + 1) % len(ps)
		a, b := ps[i], ps[j]
		if cross(a, b, point) >= 0 {
			return false
		}
	}
	return true
}

func convexHull(points []Point) []Point {
	sort.Sort(Points(points))
	n := len(points)
	up := make([]Point, n)
	var p int
	for i := 0; i < n; i++ {
		for p >= 2 && cross(up[p-2], up[p-1], points[i]) > 0 {
			p--
		}
		up[p] = points[i]
		p++
	}

	down := make([]Point, n)
	var q int
	for i := n - 1; i >= 0; i-- {
		for q >= 2 && cross(down[q-2], down[q-1], points[i]) > 0 {
			q--
		}
		down[q] = points[i]
		q++
	}
	for i := 1; i < q-1; i++ {
		up[p] = down[i]
		p++
	}
	return up[:p]
}
