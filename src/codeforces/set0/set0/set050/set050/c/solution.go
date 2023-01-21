package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	ps := make([][]int, n)
	for i := 0; i < n; i++ {
		ps[i] = readNNums(reader, 2)
	}

	res := solve(ps)

	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(ps [][]int) int {

	pts := make([]pt, len(ps))

	for i := 0; i < len(ps); i++ {
		pts[i] = pt{ps[i][0], ps[i][1]}
	}

	hull := convex_hull(pts)

	var res int

	moves := func(a, b pt) int {
		dx := abs(a.x - b.x)
		dy := abs(a.y - b.y)
		// x := min(dx, dy)
		return max(dx, dy)
	}

	for i := 0; i < len(hull); i++ {
		j := (i + 1) % len(hull)
		res += moves(hull[i], hull[j])
	}

	return res + 4
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type pt struct {
	x int
	y int
}

func orientation(a, b, c pt) int {
	v := int64(a.x)*int64(b.y-c.y) + int64(b.x)*int64(c.y-a.y) + int64(c.x)*int64(a.y-b.y)
	if v < 0 {
		return -1
	}
	if v > 0 {
		return 1
	}
	return 0
}

func cw(a, b, c pt, include_collinear bool) bool {
	v := orientation(a, b, c)
	return v < 0 || include_collinear && v == 0
}

func collinear(a, b, c pt) bool {
	return orientation(a, b, c) == 0
}

func distance(a, b pt) int64 {
	dx := int64(a.x - b.x)
	dy := int64(a.y - b.y)
	return dx*dx + dy*dy
}

func convex_hull(pts []pt) []pt {

	p0 := pts[0]

	for i := 1; i < len(pts); i++ {
		if pts[i].y < p0.y || pts[i].y == p0.y && pts[i].x < p0.x {
			p0 = pts[i]
		}
	}

	sort.Slice(pts, func(i, j int) bool {
		v := orientation(p0, pts[i], pts[j])
		if v == 0 {
			return distance(p0, pts[i]) < distance(p0, pts[j])
		}
		return v < 0
	})

	n := len(pts)

	res := make([]pt, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 1 && !cw(res[p-2], res[p-1], pts[i], false) {
			p--
		}
		res[p] = pts[i]
		p++
	}

	return res[:p]
}

func reverse(arr []pt) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
