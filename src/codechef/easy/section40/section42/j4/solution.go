package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		readString(reader)
		n := readNum(reader)
		X, Y := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			X[i], Y[i] = readTwoNums(reader)
		}
		R, C := solve(X, Y)

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%0.2f %d\n", R[i], C[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(X []int, Y []int) ([]float64, []int) {
	n := len(X)
	pts := make([]Point, n)
	for i := 0; i < n; i++ {
		pts[i] = Point{X[i], Y[i], i}
	}

	sort.Sort(Points(pts))

	// for r * r
	radius := make([]float64, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		var rad int64 = 1 << 60
		if i > 0 {
			rad = Distance(pts[i], pts[i-1])
		}
		if i+1 < n {
			rad = min(rad, Distance(pts[i], pts[i+1]))
		}
		for j := i - 2; j >= 0; j-- {
			d := Distance(pts[j], pts[i])
			if d > rad {
				break
			}
			rad = d
		}

		for j := i + 1; j < n; j++ {
			d := Distance(pts[j], pts[i])
			if d > rad {
				break
			}
			rad = d
		}
		radius[pts[i].id] = math.Sqrt(float64(rad))
		rad *= 4
		for j := i - 1; j >= 0; j-- {
			d := Distance(pts[j], pts[i])
			if d > rad {
				break
			}
			cnt[pts[i].id]++
		}
		for j := i + 1; j < n; j++ {
			d := Distance(pts[j], pts[i])
			if d > rad {
				break
			}
			cnt[pts[i].id]++
		}
	}
	return radius, cnt
}

type Point struct {
	x  int
	y  int
	id int
}

func Distance(a, b Point) int64 {
	dx := b.x - a.x
	dy := b.y - a.y
	return int64(dx)*int64(dx) + int64(dy)*int64(dy)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Points []Point

func (this Points) Len() int {
	return len(this)
}

func (this Points) Less(i, j int) bool {
	return this[i].x < this[j].x || this[i].x == this[j].x && this[i].y < this[j].y
}

func (this Points) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
