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
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		res := solve(n, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, P [][]int) int64 {
	pp := make([]Point, n)
	for i := 0; i < n; i++ {
		pp[i] = Point{P[i][0], P[i][1]}
	}
	a := find(pp)
	for i := 0; i < n; i++ {
		pp[i] = Point{P[i][1], P[i][0]}
	}
	b := find(pp)
	if a < b {
		return a
	}
	return b
}

func find(p []Point) int64 {
	sort.Sort(BX(p))
	n := len(p)
	sufm := make([]int, n)
	sufx := make([]int, n)
	sufm[n-1] = p[n-1].y
	sufx[n-1] = p[n-1].y
	for i := n - 2; i >= 0; i-- {
		sufm[i] = min(sufm[i+1], p[i].y)
		sufx[i] = max(sufx[i+1], p[i].y)
	}
	best := int64(p[n-1].x-p[0].x) * int64(sufx[0]-sufm[0])

	m, x := p[0].y, p[0].y

	for i := 1; i < n; i++ {
		l := int64(x-m) * int64(p[i-1].x-p[0].x)
		l += int64(sufx[i]-sufm[i]) * int64(p[n-1].x-p[i].x)
		if l < best {
			best = l
		}
		m = min(m, p[i].y)
		x = max(x, p[i].y)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Point struct {
	x, y int
}

type BX []Point

func (this BX) Len() int {
	return len(this)
}

func (this BX) Less(i, j int) bool {
	return this[i].x < this[j].x
}

func (this BX) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
