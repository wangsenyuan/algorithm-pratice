package main

import (
	"bufio"
	"fmt"
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
		n, c := readTwoNums(reader)
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		a, b := solve(n, c, P)

		fmt.Printf("%d %d\n", a, b)
	}
}

func solve(n int, c int, P [][]int) (int, int) {
	pts := make([]Point, n)

	for i := 0; i < n; i++ {
		pts[i] = Point{P[i][0], P[i][1], (P[i][0]%c + c) % c}
	}

	sort.Sort(Points(pts))

	var cnt, dist int

	for i, j := 0, 0; i < n; i = j {
		median := i - 1
		for j < n && pts[j].y-pts[j].x == pts[i].y-pts[i].x && pts[j].c == pts[i].c {
			j++
			if (j-i)%2 == 1 {
				median++
			}
		}

		cnt++
		for k := i; k < j; k++ {
			dist += abs(pts[k].x-pts[median].x) / c
		}
	}

	return cnt, dist
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Point struct {
	x, y int
	c    int
}

type Points []Point

func (pts Points) Len() int {
	return len(pts)
}

func (pts Points) Less(i, j int) bool {
	a, b := pts[i], pts[j]
	if a.y-a.x != b.y-b.x {
		return a.y-a.x < b.y-b.x
	}

	if a.c != b.c {
		return a.c < b.c
	}

	return a.x < b.x
}

func (pts Points) Swap(i, j int) {
	pts[i], pts[j] = pts[j], pts[i]
}
