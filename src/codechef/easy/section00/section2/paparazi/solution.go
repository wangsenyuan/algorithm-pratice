package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		H := readNNums(reader, n)
		res := solve(n, H)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, H []int) int {
	vecs := make([]Point, n)
	vecs[0] = Point{H[0], 0}
	vecs[1] = Point{H[1], 1}
	var ans = 1
	var p = 2
	for i := 2; i < n; i++ {
		cur := Point{H[i], i}
		for p > 1 && (cur.Sub(vecs[p-2])).Cross(vecs[p-1].Sub(vecs[p-2])) >= 0 {
			p--
		}
		ans = max(ans, i-vecs[p-1].y)
		vecs[p] = cur
		p++
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
	x, y int
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y}
}

func (this Point) Cross(that Point) int {
	res := int64(this.x)*int64(that.y) - int64(that.x)*int64(this.y)
	if res < 0 {
		return -1
	}
	if res > 0 {
		return 1
	}
	return 0
}
