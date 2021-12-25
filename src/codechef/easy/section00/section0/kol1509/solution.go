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
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readOneNum(scanner)
	for i := 0; i < t; i++ {
		n := readOneNum(scanner)
		nums := readNNums(scanner, n)
		fmt.Println(solve(n, nums))
	}
}

func solve(n int, xs []int) int {
	points := make([]Point, n*n/2+1)
	origin := -1
	p := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			point := Point{xs[i], xs[j]}
			points[p] = point
			p++
			if origin < 0 {
				origin = i
				continue
			}
			tmp := points[origin]
			if tmp.y > point.y || tmp.y == point.y && tmp.x > point.x {
				origin = i
			}
		}
	}
	points = points[:p]
	points[0], points[origin] = points[origin], points[0]
	points = sortPoints(points, points[0])
	m := 2
	for i := 2; i < p; i++ {
		for m > 1 && cross(points[m-2], points[m-1], points[i]) <= 0 {
			m--
		}
		points[m] = points[i]
		m++
	}
	var area int

	for i := 1; i <= m; i++ {
		x1, y1 := points[i-1].x, points[i-1].y
		x2, y2 := points[i%m].x, points[i%m].y
		tmp := x1*y2 - x2*y1
		area += tmp
	}

	if area < 0 {
		return -area
	}
	return area
}

func sortPoints(ps []Point, origin Point) []Point {
	points := Points{origin, ps[1:]}
	sort.Sort(points)
	return ps
}

type Point struct {
	x int
	y int
}

// == 0, p1, p2, p3 collinear
// > 0, p1p3 turn left with p1p2
// < 0, p1p3 turn right with p1p2
func cross(p1, p2, p3 Point) int {
	x1, y1 := p1.x, p1.y
	x2, y2 := p2.x, p2.y
	x3, y3 := p3.x, p3.y
	return (x2-x1)*(y3-y1) - (y2-y1)*(x3-x1)
}

func squareDist(p1, p2 Point) int {
	x1, y1 := p1.x, p1.y
	x2, y2 := p2.x, p2.y
	a := x2 - x1
	b := y2 - y1
	return a*a + b*b
}

type Points struct {
	origin Point
	points []Point
}

func (this Points) Len() int {
	return len(this.points)
}

func (this Points) Less(i, j int) bool {
	p1 := this.origin
	p2 := this.points[i]
	p3 := this.points[j]
	res := cross(p1, p2, p3)
	if res > 0 {
		//p3 comes after p2
		return true
	} else if res < 0 {
		//p3 comes before p2
		return false
	}
	// short comes first
	return squareDist(p1, p2) < squareDist(p1, p3)
}

func (this Points) Swap(i, j int) {
	this.points[i], this.points[j] = this.points[j], this.points[i]
}
