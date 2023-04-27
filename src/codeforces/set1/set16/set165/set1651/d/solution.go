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
		pts := make([][]int, n)
		for i := 0; i < n; i++ {
			pts[i] = readNNums(reader, 2)
		}
		res := solve(pts)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
		}
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(points [][]int) [][]int {
	n := len(points)
	pts := make(map[Point]bool)
	for i := 0; i < n; i++ {
		pts[Point{points[i][0], points[i][1]}] = true
	}

	que := make([]Point, 0, 1)
	ans := make(map[Point]Point)

	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		cur := Point{x, y}
		for j := 0; j < 4; j++ {
			u, v := x+dd[j], y+dd[j+1]
			p := Point{u, v}
			if !pts[p] {
				que = append(que, cur)
				ans[cur] = p
			}
		}
	}

	for i := 0; i < len(que); i++ {
		cur := que[i]
		x, y := cur.x, cur.y
		for j := 0; j < 4; j++ {
			u, v := x+dd[j], y+dd[j+1]
			p := Point{u, v}
			if _, ok := ans[p]; ok || !pts[p] {
				continue
			}
			ans[p] = ans[cur]
			que = append(que, p)
		}
	}

	var res [][]int

	for _, p := range points {
		x, y := p[0], p[1]
		cur := ans[Point{x, y}]
		res = append(res, []int{cur.x, cur.y})
	}
	return res
}

type Point struct {
	x int
	y int
}
