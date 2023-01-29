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
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 2)
	}
	res := solve(points)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(points [][]int) int {
	n := len(points)

	type Point struct {
		x  int
		y  int
		id int
	}

	P := make([]Point, n)

	for i := 0; i < n; i++ {
		P[i] = Point{points[i][0], points[i][1], i}
	}

	X := make(map[int][]Point)
	Y := make(map[int][]Point)

	add := func(m map[int][]Point, k int, p Point) {
		if _, ok := m[k]; !ok {
			m[k] = make([]Point, 0, 1)
		}
		m[k] = append(m[k], p)
	}

	for _, p := range P {
		add(X, p.x, p)
		add(Y, p.y, p)
	}

	flag := make([]int, n)

	for _, vs := range X {
		// same x
		sort.Slice(vs, func(i, j int) bool {
			return vs[i].y < vs[j].y
		})

		for i := 1; i+1 < len(vs); i++ {
			if vs[i].y > vs[0].y && vs[i].y < vs[len(vs)-1].y {
				flag[vs[i].id]++
			}
		}
	}
	for _, vs := range Y {
		// same y
		sort.Slice(vs, func(i, j int) bool {
			return vs[i].x < vs[j].x
		})

		for i := 1; i+1 < len(vs); i++ {
			if vs[i].x > vs[0].x && vs[i].x < vs[len(vs)-1].x {
				flag[vs[i].id]++
			}
		}
	}
	var res int

	for i := 0; i < n; i++ {
		if flag[i] == 2 {
			res++
		}
	}
	return res
}
