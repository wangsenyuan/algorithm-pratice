package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	pts := readNNums(reader, 6)

	res := solve(pts)

	fmt.Println(res)
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

var d = []int{-1, 0, 1, 0, -1}

func solve(pts []int) string {
	ok := checkRight(pts)
	if ok {
		return "RIGHT"
	}

	for i := 0; i < 3; i++ {
		x, y := pts[i*2], pts[i*2+1]

		for k := 0; k < 4; k++ {
			pts[i*2] = x + d[k]
			pts[i*2+1] = y + d[k+1]
			if checkRight(pts) {
				return "ALMOST"
			}
		}

		pts[i*2] = x
		pts[i*2+1] = y
	}

	return "NEITHER"
}

func checkRight(pts []int) bool {
	if pts[0] == pts[2] && pts[2] == pts[4] {
		return false
	}

	if pts[1] == pts[3] && pts[3] == pts[5] {
		return false
	}

	a := distance(pts[0:2], pts[2:4])
	b := distance(pts[0:2], pts[4:6])
	c := distance(pts[2:4], pts[4:6])

	if a == 0 || b == 0 || c == 0 {
		return false
	}

	return a+b == c || a+c == b || b+c == a
}

func distance(a, b []int) int {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	return dx*dx + dy*dy
}
