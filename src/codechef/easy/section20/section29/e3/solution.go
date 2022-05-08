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
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, 2*n)
		res := solve(n, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, P []int) int {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		points[i] = NewPoint(P[2*i], P[2*i+1])
	}

	sort.Sort(Points(points))

	hull := make([]Point, 2*n)
	var pos int

	hull[pos] = points[0]
	pos++

	for i := 1; i < n; i++ {
		cur := points[i]
		for pos >= 2 && (cur.Sub(hull[pos-2]).Cross(cur.Sub(hull[pos-1]))) >= 0 {
			pos--
		}
		hull[pos] = cur
		pos++
	}

	start := pos - 1

	for i := n - 2; i >= 0; i-- {
		cur := points[i]
		for pos-start >= 2 && (cur.Sub(hull[pos-2]).Cross(cur.Sub(hull[pos-1]))) >= 0 {
			pos--
		}
		hull[pos] = cur
		pos++
	}
	pos--

	var ans int

	for i := 0; i < pos; i++ {
		j := (i + 1) % pos
		k := (j + 1) % pos

		for k != i && area(hull[i], hull[j], hull[k]) < area(hull[i], hull[j], hull[(k+1)%pos]) {
			k = (k + 1) % pos
		}
		if k == i {
			continue
		}
		kk := (k + 1) % pos
		for j != kk && k != i {
			ans = max(ans, area(hull[i], hull[j], hull[k]))

			for k != i && area(hull[i], hull[j], hull[k]) < area(hull[i], hull[j], hull[(k+1)%pos]) {
				k = (k + 1) % pos
			}

			j = (j + 1) % pos
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (this Point) Less(that Point) bool {
	return this.x < that.x || this.x == that.x && this.y < that.y
}

func (this Point) Sub(that Point) Point {
	return NewPoint(this.x-that.x, this.y-that.y)
}

func (this Point) Cross(that Point) int {
	return this.x*that.y - this.y*that.x
}

func area(a, b, c Point) int {
	return abs(b.Sub(a).Cross(c.Sub(a)))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Points []Point

func (this Points) Len() int {
	return len(this)
}

func (this Points) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Points) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
