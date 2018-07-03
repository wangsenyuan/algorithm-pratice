package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(scanner, 2)
		}
		res := solve(n, points)
		for _, ans := range res {
			fmt.Printf("%d %d\n", ans[0], ans[1])
		}
		tc--
	}
}

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

func solve(n int, points [][]int) [][]int {
	ps := make(Points, n)
	for i := 0; i < n; i++ {
		ps[i] = Point{points[i][0], points[i][1]}
	}
	sort.Sort(ps)

	findHull := func(lower bool) Points {
		sign := 1
		if lower {
			sign = -1
		}
		res := make(Points, n)
		var idx int
		for i := 0; i < n; i++ {
			for idx >= 2 && cross(res[idx-2], res[idx-1], ps[i])*sign >= 0 {
				idx--
			}
			res[idx] = ps[i]
			idx++
		}
		return res[:idx]
	}

	upper := findHull(false)
	lower := findHull(true)

	checkGood := func(hull Points, p Point, lower bool) bool {
		sign := 1
		if !lower {
			sign = -1
		}
		it := sort.Search(len(hull), func(i int) bool {
			return hull[i].x >= p.x
		})

		if it == len(hull) || it == 0 {
			return false
		}

		res := sign * cross(hull[it-1], hull[it], p)
		return res > 0
	}

	inside := func(p Point) bool {
		return checkGood(lower, p, true) && checkGood(upper, p, false)
	}

	var p0 Point

	for i := 2; i < n; i++ {
		p0.x = (ps[0].x + ps[i].x) / 2
		p0.y = (ps[0].y + ps[i].y) / 2
		if inside(p0) {
			break
		}
	}

	que := make([]Point, n/10+4)
	front, end := 0, 0
	que[end] = p0
	end++
	checked := make(map[Point]bool)
	checked[p0] = true
	for end < n/10 && front < end {
		u := que[front]
		front++

		vs := []Point{
			Point{u.x - 1, u.y},
			Point{u.x + 1, u.y},
			Point{u.x, u.y - 1},
			Point{u.x, u.y + 1},
		}

		for _, v := range vs {
			if !checked[v] && inside(v) {
				checked[v] = true
				que[end] = v
				end++
			}
		}

	}
	res := make([][]int, n/10)
	for i := 0; i < n/10; i++ {
		res[i] = []int{que[i].x, que[i].y}
	}
	return res
}

func solve1(n int, points [][]int) [][]int {
	set := make(map[Point]bool)

	for i := 0; len(set) < n/10; i += 9 {
		for j := i; j < i+9 && j < n; j++ {
			for k := j + 2; k < i+9 && k < n; k++ {
				a, b := points[j], points[k]
				if (a[0]+b[0])%2 == 0 && (a[1]+b[1])%2 == 0 {
					set[Point{(a[0] + b[0]) / 2, (a[1] + b[1]) / 2}] = true
				}
			}
		}
	}

	res := make([][]int, len(set))
	var idx int
	for p := range set {
		res[idx] = []int{p.x, p.y}
		idx++
	}
	return res[:n/10]
}

type Point struct {
	x, y int
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y}
}

func cross(a, b, c Point) int {
	ab := b.Sub(a)
	bc := c.Sub(b)
	x := int64(ab.x) * int64(bc.y)
	y := int64(ab.y) * int64(bc.x)
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

type Points []Point

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	return ps[i].x < ps[j].x || (ps[i].x == ps[j].x && ps[i].y < ps[j].y)
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
