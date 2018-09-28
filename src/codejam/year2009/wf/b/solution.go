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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for x := 1; x <= tc; x++ {
		n := readNum(scanner)
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(scanner, 2)
		}
		res := solve(n, points)
		fmt.Printf("Case #%d: %.12f\n", x, res)
	}
}

const BILLION = 1000000000

const INF = 1e20

func solve(n int, points [][]int) float64 {
	px := make([]Point, n)
	for i := 0; i < n; i++ {
		px[i] = Point{2*points[i][0] - BILLION, 2*points[i][1] - BILLION}
	}
	py := make([]Point, n)
	copy(py, px)
	sort.Sort(PointsX(px))
	sort.Sort(PointsY(py))

	return calc(n, px, py) / 2
}

func calc(n int, px []Point, py []Point) float64 {
	if n < 3 {
		return INF
	}

	half := n / 2
	left, right := px[:half], px[half:]

	pyl := make([]Point, 0, half)
	pyr := make([]Point, 0, n-half)

	mid := Point{(px[half-1].x + px[half].x) / 2, (px[half-1].y + px[half].y) / 2}

	for i := 0; i < n; i++ {
		p := py[i]
		if compareX(p, mid) {
			pyl = append(pyl, p)
		} else {
			pyr = append(pyr, p)
		}
	}
	res := INF
	res = math.Min(res, calc(half, left, pyl))
	res = math.Min(res, calc(n-half, right, pyr))

	var margin float64
	if res > INF/2 {
		margin = 2 * BILLION
	} else {
		margin = res / 2
	}

	closeToTheLine := make([]Point, 0, n)

	for i, j := 0, 0; i < n; i++ {
		p := py[i]
		if absDiff2(p.x, mid.x) > margin {
			continue
		}
		for j < len(closeToTheLine) && absDiff2(p.y, closeToTheLine[j].y) > margin {
			j++
		}

		for k := j; k < len(closeToTheLine); k++ {
			for l := k + 1; l < len(closeToTheLine); l++ {
				a, b := closeToTheLine[k], closeToTheLine[l]
				res = math.Min(res, perimeter(p, a, b))
			}
		}
		closeToTheLine = append(closeToTheLine, p)
	}

	return res
}

func absDiff(a int, b float64) float64 {
	return math.Abs(float64(a) - b)
}

func absDiff2(a int, b int) float64 {
	return math.Abs(float64(a) - float64(b))
}

func perimeter(a, b, c Point) float64 {
	return dist(a, b) + dist(b, c) + dist(a, c)
}

func dist(a, b Point) float64 {
	return math.Sqrt(sqr(a.x-b.x) + sqr(a.y-b.y))
}

func sqr(x int) float64 {
	return float64(x) * float64(x)
}

type Point struct {
	x, y int
}

type PointsX []Point

func (this PointsX) Len() int {
	return len(this)
}

func (this PointsX) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this PointsX) Less(i, j int) bool {
	return compareX(this[i], this[j])
}

type PointsY []Point

func (this PointsY) Len() int {
	return len(this)
}

func (this PointsY) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this PointsY) Less(i, j int) bool {
	return compareY(this[i], this[j])
}

func compareX(a, b Point) bool {
	return a.x < b.x || a.x == b.x && a.y < b.y
}

func compareY(a, b Point) bool {
	return a.y < b.y || a.y == b.y && a.x < b.x
}
