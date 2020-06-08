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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(reader, 2)
		}
		fmt.Println(solve(n, points))
	}
}

func solve(n int, points [][]int) int64 {
	ps := make([]Point, n)
	for i := 0; i < n; i++ {
		ps[i] = Point{points[i][0], points[i][1], i}
	}

	outer := convexHull(ps)

	inOuter := make([]bool, n)

	for i := 0; i < len(outer); i++ {
		inOuter[outer[i].i] = true
	}

	var m int

	for i := 0; i < n; i++ {
		if inOuter[ps[i].i] {
			continue
		}
		ps[m] = ps[i]
		m++
	}

	inner := convexHull(ps[:m])

	if len(inner) == 0 {
		return -1
	}

	var area int64

	for i := 1; i+1 < len(outer); i++ {
		area += abs(triangleArea(outer[0], outer[i], outer[i+1]))
	}

	var mn int64 = math.MaxInt64
	var p int
	for i := 0; i < len(outer); i++ {
		cur := abs(triangleArea(outer[i], outer[(i+1)%len(outer)], inner[p]))

		var start = p

		for {
			nx := abs(triangleArea(outer[i], outer[(i+1)%len(outer)], inner[(p+1)%len(inner)]))
			if nx <= cur {
				cur = nx
				p = (p + 1) % len(inner)
			} else {
				break
			}
			if start == p {
				break
			}
		}

		mn = min(mn, cur)
	}

	return area - mn
}

type Point struct {
	first, second int
	i             int
}

type Points []Point

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Less(i, j int) bool {
	a, b := ps[i], ps[j]
	return a.first < b.first || a.first == b.first && a.second < b.second
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func convexHull(points []Point) []Point {
	n := len(points)
	if n <= 1 {
		return points
	}

	sort.Sort(Points(points))

	var k int

	q := make([]Point, 2*n)

	for i := 0; i < n; i++ {
		for k >= 2 && !cw(q[k-2], q[k-1], points[i]) {
			k--
		}
		q[k] = points[i]
		k++
	}
	t := k
	for i := n - 2; i >= 0; i-- {
		for k > t && !cw(q[k-2], q[k-1], points[i]) {
			k--
		}
		q[k] = points[i]
		k++
	}

	k--
	if q[0].i == q[1].i {
		k--
	}
	return q[:k]
}

func cw(a, b, c Point) bool {
	return (b.first-a.first)*(c.second-a.second)-(b.second-a.second)*(c.first-a.first) < 0
}

func triangleArea(p1, p2, p3 Point) int64 {
	return int64(p2.first-p1.first)*int64(p3.second-p1.second) - int64(p2.second-p1.second)*int64(p3.first-p1.first)
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
