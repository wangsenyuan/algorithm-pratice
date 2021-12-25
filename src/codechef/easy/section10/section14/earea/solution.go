package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := 1
	for tc > 0 {
		tc--
		n := readNum(reader)
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		res := solve(n, P)
		buf.WriteString(fmt.Sprintf("%.8f\n", res))
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

func solve(n int, P [][]int) float64 {
	var a int64
	var ap int64
	pp := make([]Point, n)
	for i := 0; i < n; i++ {
		pp[i] = Point{int64(P[i][0]), int64(P[i][1])}
	}

	for i := 0; i < n; i++ {
		p := pp[i].Sub(pp[(i+n-1)%n])
		q := pp[(i+1)%n].Sub(pp[i])
		z := p.Cross(q)
		a += z
		ap += p.Cross(pp[i])
	}
	if ap < 0 {
		ap = -ap
	}
	return float64(4*ap-a) / 8
}

type Point struct {
	x int64
	y int64
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y}
}

func (this Point) Add(that Point) Point {
	return Point{this.x + that.x, this.y + that.y}
}

func (this Point) Cross(that Point) int64 {
	return this.x*that.y - this.y*that.x
}
