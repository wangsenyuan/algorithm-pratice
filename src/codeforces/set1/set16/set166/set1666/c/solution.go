package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	points := make([][]int, 3)
	for i := 0; i < 3; i++ {
		points[i] = readNNums(reader, 2)
	}
	res := solve(points)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, seg := range res {
		for i := 0; i < len(seg); i++ {
			buf.WriteString(fmt.Sprintf("%d ", seg[i]))
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
func solve(points [][]int) [][]int {
	slices.SortFunc(points, func(a, b []int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	var res [][]int
	// 保证垂直方向
	a, b, c := points[0], points[1], points[2]
	if a[0] <= b[0] && b[0] <= c[0] || a[0] >= b[0] && b[0] >= c[0] {
		// b在(a, c)的中间
		res = append(res, []int{a[0], a[1], a[0], b[1]})
		res = append(res, []int{a[0], b[1], c[0], b[1]})
		res = append(res, []int{c[0], b[1], c[0], c[1]})
		return res
	}

	if a[0] <= c[0] {
		if b[0] < a[0] {
			// b 在a的前面
			res = append(res, []int{a[0], a[1], a[0], b[1]})
			res = append(res, []int{a[0], b[1], b[0], b[1]})
			res = append(res, []int{a[0], b[1], a[0], c[1]})
			res = append(res, []int{a[0], c[1], c[0], c[1]})
			return res
		}
		// b[0] > c[0], b在c的后面
		res = append(res, []int{a[0], a[1], c[0], a[1]})
		res = append(res, []int{c[0], a[1], c[0], b[1]})
		res = append(res, []int{c[0], b[1], b[0], b[1]})
		res = append(res, []int{c[0], b[1], c[0], c[1]})
		return res
	}
	// a[0] > c[0]
	if b[0] < c[0] {
		// b在c的前面
		res = append(res, []int{b[0], b[1], c[0], b[1]})
		res = append(res, []int{c[0], b[1], c[0], c[1]})
		res = append(res, []int{c[0], b[1], a[0], b[1]})
		res = append(res, []int{a[0], b[1], a[0], a[1]})
		return res
	}
	// b[0] > a[0], b在a的后面
	res = append(res, []int{c[0], c[1], a[0], c[1]})
	res = append(res, []int{a[0], c[1], a[0], b[1]})
	res = append(res, []int{a[0], b[1], b[0], b[1]})
	res = append(res, []int{a[0], b[1], a[0], a[1]})
	return res
}
