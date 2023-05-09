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
	pts := make([][]int, n)
	for i := 0; i < n; i++ {
		pts[i] = readNNums(reader, 2)
	}

	res := solve(pts)

	fmt.Printf("%.1f\n", res)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) halfplane() int {
	if p.y > 0 || p.y == 0 && p.x > 0 {
		return 1
	}
	return 0
}

func (a Point) cross(b Point) int64 {
	first := int64(a.x) * int64(b.y)
	second := int64(b.x) * int64(a.y)
	return first - second
}

func (a Point) less(b Point) bool {
	return a.x < b.x || a.x == b.x && a.y < b.y
}

func (a Point) add(b Point) Point {
	return NewPoint(a.x+b.x, a.y+b.y)
}

func (a Point) minus() Point {
	return NewPoint(-a.x, -a.y)
}

func solve(pts [][]int) float64 {
	n := len(pts)
	// n <= 150
	A := make([]Point, n)
	for i := 0; i < n; i++ {
		A[i] = NewPoint(pts[i][0], pts[i][1])
	}

	lessAngle := func(a, b Point) bool {
		ah := a.halfplane()
		bh := b.halfplane()
		if ah != bh {
			return ah > bh
		}
		return a.cross(b) > 0
	}

	sort.Slice(A, func(i, j int) bool {
		return lessAngle(A[i], A[j])
	})

	type Pair struct {
		first  Point
		second int64
	}

	var B []Pair

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			aij := A[i].cross(A[j])
			if aij < 0 {
				break
			}
			B = append(B, Pair{A[i].add(A[j]), aij})
			for k := j; k < n; k++ {
				if A[j].cross(A[k]) < 0 {
					break
				}
				sum := A[i].add(A[j]).add(A[k])
				if A[i].cross(sum) >= 0 {
					B = append(B, Pair{sum, aij + abs(A[k].cross(sum))})
				}
			}
		}
	}
	sort.Slice(B, func(i, j int) bool {
		if B[i].first.less(B[j].first) {
			return true
		}
		if B[i].first == B[j].first {
			return B[i].second < B[j].second
		}
		return false
	})
	ln := 0
	for i := 0; i < len(B); {
		j := i
		for i < len(B) && B[i].first == B[j].first {
			i++
		}
		B[ln] = B[i-1]
		ln++
	}
	B = B[:ln]

	var ans int64

	for i := 0; i < ln; i++ {
		sum := B[i].first.minus()
		j := search(ln, func(j int) bool {
			return sum.less(B[j].first) || sum == B[j].first
		})
		if j < ln && B[j].first == sum {
			ans = max(ans, B[i].second+B[j].second)
		}
	}
	return float64(ans) / 2
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) >> 1
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
